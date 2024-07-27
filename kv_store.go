package hermeskv

import (
	"log"
	"sync"
	"time"
)

type ValMetaDeta struct {
	value       interface{}
	isTombStone bool
}
type Store struct {
	// store the key value pairs
	sync.RWMutex
	wg          *sync.WaitGroup
	shutdown    chan struct{}
	globalState map[string]*ValMetaDeta
	FIFO        *DoublyLinkedList
	LRU         *DoublyLinkedList
	capacity    int
}

//go:generate mockgen -destination=./mocks/mock_store.go -package=mocks hermeskv StoreIface
type StoreIface interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	Close()
}

var _ StoreIface = (*Store)(nil)

func GetStore(capacity int) *Store {
	if capacity <= 0 {
		panic("cache capacity cant be zero")
	}

	// map[key] = val
	// key = string
	// val = nodeRef stored in the DLL
	s := &Store{
		globalState: make(map[string]*ValMetaDeta),
		FIFO:        getDLL(),
		LRU:         getDLL(),
		capacity:    capacity,
		RWMutex:     sync.RWMutex{},
		shutdown:    make(chan struct{}),
		wg:          &sync.WaitGroup{},
	}

	// start the background purger
	s.wg.Add(1)
	interval := 5 * time.Second
	go s.purger(interval)

	return s
}

// clean up old entries by checking the capacity has been breached at regular intervals when some TTL is set on the keys.
// there is a possibility that we can end up adding more entries than the current capacity if the purger hasn't cleaned up yet
// but thats fine cause eventually they will be deleted.

// well no, eventual consistent seems to be problem here cause the in-memory operations(set,get,delete) are quite fast.
// we can't guarantee when and if the purger goroutine runs and deletes the head node which leads to inconsistent results.

// this is similar to redis's background thread evicting the entries according to the set policy.
// but the only hard limit for redis is the underlying RAM so redis can use more RAM than available but EVENTUALLY it will clean up the old entries.

func (s *Store) purger(interval time.Duration) {
	newTicker := time.NewTicker(interval)

	for {
		select {
		case <-newTicker.C:
			for s.LRU.capacity > s.capacity {
				s.RWMutex.Lock()
				defer s.RWMutex.Unlock()

				log.Println("capacity breached. deleting the head node")
				node := s.LRU.deleteHead()
				delete(s.globalState, node.key)
			}
		case <-s.shutdown:
			newTicker.Stop()
			s.wg.Done()
			return

		}
	}

}

func (s *Store) Set(key string, value interface{}) error {
	// take writer lock while adding to the cache
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// check the cur len > capacity, delete the head node.
	for s.LRU.capacity >= s.capacity {
		// evict the head node and update the capacity
		node := s.LRU.deleteHead()
		delete(s.globalState, node.key)
	}

	// add node to the DLL
	node := getNode(key, value, nil, nil)
	newNode := s.LRU.addNode(node)
	s.globalState[key] = &ValMetaDeta{
		// for globalState, value is the node reference in the doubly linked list.
		value:       newNode,
		isTombStone: false,
	}

	return nil
}

// for LRU, whenever a key is referenced/updated, move that key from front to the back of the DLL.
func (s *Store) Get(key string) (interface{}, error) {
	// shared reader lock for accessing the cache
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	valMeta, ok := s.globalState[key]

	if !ok || valMeta.isTombStone {
		return nil, ErrNoKey
	}

	node := valMeta.value.(*Node)

	// update the DLL
	s.LRU.deleteNode(node)
	s.LRU.addNode(node)

	return node.val, nil
}

func (s *Store) Delete(key string) error {
	// take writer lock while deleting from the cache
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	valMeta, ok := s.globalState[key]
	if !ok || valMeta.isTombStone {
		return ErrNoKey
	}
	node := valMeta.value.(*Node)

	// for the globalState, we aren't marking it as a tombStone but actually deleting it.
	delete(s.globalState, key)
	s.LRU.deleteNode(node)

	return nil
}

// equivalent to closing a connection. useful in graceful shutdowns
func (s *Store) Close() {
	close(s.shutdown)
	s.wg.Wait()
}
