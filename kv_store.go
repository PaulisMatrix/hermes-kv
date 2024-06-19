package hermeskv

import (
	"sync"
)

type ValMetaDeta struct {
	value       interface{}
	isTombStone bool
}
type Store struct {
	// store the key value pairs
	sync.RWMutex
	globalState map[string]*ValMetaDeta
	FIFO        *DoublyLinkedList
	capacity    int
	StoreIface
}

type StoreIface interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
}

var _ StoreIface = (*Store)(nil)

func GetNewKV(capacity int) *Store {
	if capacity <= 0 {
		panic("cache capacity cant be zero")
	}

	// map[key] = val
	// key = string
	// val = nodeRef stored in the DLL
	s := &Store{
		globalState: make(map[string]*ValMetaDeta),
		FIFO:        getDLL(),
		capacity:    capacity,
		RWMutex:     sync.RWMutex{},
	}
	return s
}

func (s *Store) Set(key string, value interface{}) error {
	// take writer lock while adding to the cache
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	// check the cur len > capacity, delete the head node.
	if s.FIFO.capacity >= s.capacity {
		// TODO: can we do this in a background thread?
		// evict the head node and update the capacity
		node := s.FIFO.deleteHead()
		delete(s.globalState, node.key)
	}

	// add node to the DLL
	node := getNode(key, value, nil, nil)
	newNode := s.FIFO.addNode(node)
	s.globalState[key] = &ValMetaDeta{
		// for globalState, value is the node reference in the doubly linked list.
		value:       newNode,
		isTombStone: false,
	}

	return nil
}

func (s *Store) Get(key string) (interface{}, error) {
	// shared reader lock for accessing the cache
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	valMeta, ok := s.globalState[key]

	if !ok || valMeta.isTombStone {
		return nil, ErrNoKey
	}

	node := valMeta.value.(*Node)
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
	s.FIFO.deleteNode(node)

	return nil
}
