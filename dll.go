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
	// head and tail nodes are basically the boundaries pointing to the first and last element
	headNode *Node
	tailNode *Node
	capacity int
}

func getDLL() *DoublyLinkedList {
	dll := &DoublyLinkedList{
		headNode: getNode("headNode", -1, nil, nil),
		tailNode: getNode("tailNode", -1, nil, nil),
		capacity: 0,
	}
	// init head and tail nodes
	dll.headNode.prev = nil
	dll.headNode.next = dll.tailNode

	dll.tailNode.next = nil
	dll.tailNode.prev = dll.headNode

	return dll
}

func (dll *DoublyLinkedList) addNode(key string, val interface{}) *Node {
	newNode := getNode(key, val, nil, nil)
	dll.capacity++

	// insert in between tail and head
	prevNode := dll.tailNode.prev
	dll.tailNode.prev = newNode
	newNode.next = dll.tailNode
	newNode.prev = prevNode
	prevNode.next = newNode
	return newNode

}

func (dll *DoublyLinkedList) display() {
	curNode := dll.headNode

	for curNode != nil {
		fmt.Println("current node val: ", curNode.val)
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
func (dll *DoublyLinkedList) deleteHead() *Node {

	head := dll.headNode.next
	dll.headNode.next = head.next
	head.next.prev = dll.headNode
	return head
}

// delete a specifc node
func (dll *DoublyLinkedList) deleteNode(node *Node) {

}
