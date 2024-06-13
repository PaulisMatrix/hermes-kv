package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDLLSet(t *testing.T) {
	newDLL := getDLL()

	// add node
	newDLL.addNode("f", 100)
	newDLL.addNode("s", 101)
	newDLL.addNode("t", 102)

	// search for 102
	node, err := newDLL.getNode(102)
	require.Nil(t, err)

	assert.EqualValues(t, 102, node.val)

}

func TestDLLDeleteHead(t *testing.T) {
	newDLL := getDLL()

	// add node
	newDLL.addNode("f", 100)
	newDLL.addNode("s", 101)
	newDLL.addNode("t", 102)

	// delete the head node
	node := newDLL.deleteHead()

	assert.EqualValues(t, node.val, 100)

}
