package hermeskv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDLLSet(t *testing.T) {
	newDLL := getDLL()

	// add nodes
	// add nodes
	n1 := getNode("f", 100, nil, nil)
	n2 := getNode("s", 101, nil, nil)
	n3 := getNode("t", 102, nil, nil)

	newDLL.addNode(n1)
	newDLL.addNode(n2)
	newDLL.addNode(n3)

	// search for 102
	node, err := newDLL.getNode(102)
	require.Nil(t, err)

	assert.EqualValues(t, 102, node.val)

}

func TestDLLDeleteHead(t *testing.T) {
	newDLL := getDLL()

	// add nodes
	// add nodes
	n1 := getNode("f", 100, nil, nil)
	n2 := getNode("s", 101, nil, nil)
	n3 := getNode("t", 102, nil, nil)

	newDLL.addNode(n1)
	newDLL.addNode(n2)
	newDLL.addNode(n3)

	// delete the head node
	node := newDLL.deleteHead()

	assert.EqualValues(t, node.val, 100)
	assert.EqualValues(t, 101, newDLL.headNode.next.val)

}

func TestDLLDeleteTail(t *testing.T) {
	newDLL := getDLL()

	// add nodes
	n1 := getNode("f", 100, nil, nil)
	n2 := getNode("s", 101, nil, nil)
	n3 := getNode("t", 102, nil, nil)

	newDLL.addNode(n1)
	newDLL.addNode(n2)
	newDLL.addNode(n3)

	// delete the tail node
	node := newDLL.deleteTail()

	assert.EqualValues(t, node.val, 102)
	assert.EqualValues(t, 101, newDLL.tailNode.prev.val)

}

func TestDLLDeleteNode(t *testing.T) {
	newDLL := getDLL()

	// add nodes
	n1 := getNode("f", 100, nil, nil)
	n2 := getNode("s", 101, nil, nil)
	n3 := getNode("t", 102, nil, nil)

	newDLL.addNode(n1)
	newDLL.addNode(n2)
	newDLL.addNode(n3)

	// delete the middle node
	newDLL.deleteNode(n2)

	assert.EqualValues(t, n2.val, 101)
	assert.EqualValues(t, 102, n1.next.val)
	assert.EqualValues(t, 100, n3.prev.val)
}
