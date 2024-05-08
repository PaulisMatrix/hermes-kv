package main

import (
	"errors"
	"fmt"
)

type Node struct {
	val  interface{}
	key  string
	next *Node
	prev *Node
}

func getNode(key string, val interface{}, prev, next *Node) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: next,
		prev: prev,
	}
}

type DoublyLinkedList struct {
	curNode  *Node
	headNode *Node
	capacity int
}

func getDLL() *DoublyLinkedList {
	dll := &DoublyLinkedList{
		curNode:  nil,
		headNode: nil,
		capacity: 0,
	}

	return dll
}

func (dll *DoublyLinkedList) addNode(key string, val interface{}) *Node {
	newNode := getNode(key, val, nil, nil)

	// adding first node
	if dll.curNode == nil {
		dll.curNode = newNode
		dll.curNode.next = nil
		dll.curNode.prev = nil

		dll.headNode = dll.curNode
		return newNode
	}

	dll.curNode.next = newNode
	newNode.prev = dll.curNode
	newNode.next = nil
	dll.curNode = newNode

	// <- head -> tail ->
	// <-head<-100 -><-tail
	// <-head<-100 -><-200 300

	dll.capacity++
	return newNode

}

func (dll *DoublyLinkedList) display() {
	curNode := dll.headNode

	for curNode != nil {
		fmt.Println("current node val", curNode.val)
		curNode = curNode.next
	}

}

func (dll *DoublyLinkedList) getNode(val interface{}) (*Node, error) {
	// linear

	curNode := dll.headNode
	for curNode != nil {
		if curNode.val == val {
			return curNode, nil
		}
		curNode = curNode.next
	}
	return nil, errors.New("node not found")
}

// pop the head node everytime
func (dll *DoublyLinkedList) deleteNode() *Node {
	head := dll.headNode

	nextNode := dll.headNode.next

	dll.headNode = nextNode
	nextNode.prev = nil

	dll.capacity--
	return head
}
