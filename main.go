package main

import (
	"errors"
	"fmt"
	"os"
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
		// evict the head node and update the capacity
		fmt.Println("capacity breached, deleting head node...")
		// how to get the key
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
		return nil, errors.New("Key doesn't exist")
	}

	node := nodeRef.(*Node)

	return node.val, nil
}

func (s *Store) Delete(key string) error {
	nodeRef, ok := s.KVMap[key]
	if !ok {
		return errors.New("Key doesn't exist")
	}
	node := nodeRef.(*Node)

	// delete from the map and from the DLL
	delete(s.KVMap, key)
	s.FIFO.deleteNode(node)

	return nil
}

func main() {
	s := GetNewKV(2)

	s.Set("hello", "world")
	s.Set("first", 100)
	s.Set("second", 200)

	val, err := s.Get("first")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("value received: ", val)
}
