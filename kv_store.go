package hermeskv

import (
	"sync"
)

type Store struct {
	// store the key value pairs
	sync.RWMutex
	KVMap    map[string]interface{}
	FIFO     *DoublyLinkedList
	capacity int
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
		KVMap:    make(map[string]interface{}, capacity),
		FIFO:     getDLL(),
		capacity: capacity,
		RWMutex:  sync.RWMutex{},
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
		delete(s.KVMap, node.key)
	}

	// add node to the DLL
	node := getNode(key, value, nil, nil)
	newNode := s.FIFO.addNode(node)
	// key -> nodeRef
	s.KVMap[key] = newNode

	return nil
}

func (s *Store) Get(key string) (interface{}, error) {
	// shared reader lock for accessing the cache
	s.RWMutex.RLock()
	defer s.RWMutex.RUnlock()

	nodeRef, ok := s.KVMap[key]

	if !ok {
		return nil, ErrNoKey
	}

	node := nodeRef.(*Node)

	return node.val, nil
}

func (s *Store) Delete(key string) error {
	// take writer lock while deleting from the cache
	s.RWMutex.Lock()
	defer s.RWMutex.Unlock()

	nodeRef, ok := s.KVMap[key]
	if !ok {
		return ErrNoKey
	}
	node := nodeRef.(*Node)

	// delete from the map and from the DLL
	delete(s.KVMap, key)
	s.FIFO.deleteNode(node)

	return nil
}
