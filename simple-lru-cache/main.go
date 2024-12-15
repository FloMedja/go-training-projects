package main

import (
	"fmt"
)

const SIZE = 5

// Node implementation for the double linked list
type Node struct {
	Value string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head *Node
	Tail *Node
	Len  int
}

type Hash map[string]*Node

// TODO: integrate mutex for thread safety
type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{
		Queue: NewQueue(),
		Hash:  Hash{},
	}
}

func NewQueue() Queue {
	// head := &Node{}
	// tail := &Node{}
	// head.Right = tail
	// tail.Left = head

	// return Queue{
	// 	Head: head,
	// 	Tail: tail,
	// }
	// Create en empty Queue might be simplier
	return Queue{
		Head: nil,
		Tail: nil,
		Len:  0,
	}

}

func (c *Cache) Append(n *Node) {
	fmt.Printf("Adding value value: %s\n", n.Value)
	if c.Queue.Len == 0 {
		c.Queue.Head = n
		c.Queue.Tail = n
	} else {
		c.Queue.Head.Left = n
		n.Right = c.Queue.Head
		c.Queue.Head = n
	}
	c.Hash[n.Value] = n
	c.Queue.Len += 1

	if c.Queue.Len > SIZE {
		c.Remove(c.Queue.Tail)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head
	fmt.Printf("%d - element(s) [", q.Len)

	for node != nil {
		fmt.Printf("{%s}", node.Value)
		if node.Right != nil {
			fmt.Printf("<--->")
		}
		node = node.Right
	}

	fmt.Printf("]\n")
}

func (c *Cache) Check(str string) {
	node := &Node{}
	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{
			Value: str,
		}
	}
	c.Append(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	if n == nil {
		fmt.Printf("Node to be removed is nil value %s\n", n.Value)
		return nil
	}
	fmt.Printf("Removing node with value %s\n", n.Value)

	// We create new links
	left := n.Left
	right := n.Right
	if left != nil {
		left.Right = n.Right
	}
	if right != nil {
		right.Left = n.Left
	}
	// We handle the head deletion
	if c.Queue.Head == n {
		c.Queue.Head = right
	}
	// We handle the queue deletion
	if c.Queue.Tail == n {
		c.Queue.Tail = left
	}

	// We remove the node
	n.Left = nil
	n.Right = nil
	c.Queue.Len -= 1
	delete(c.Hash, n.Value)

	return n

}

func (c *Cache) RemoveNodeWithVal(val string) *Node {
	fmt.Printf("Removing first node with value %s\n", val)
	current := c.Queue.Head
	for current != nil {
		if current.Value == val {
			current.Left = current.Right
			current.Right = current.Left
			return current
		}
		current = current.Right
	}
	return nil
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"parrot", "orange", "dragonfruit", "cat", "dog", "orange", "mangue", "pêche", "pêche", "pêche", "pêche", "pêche", "orange"} {
		cache.Check(word)
		cache.Display()
	}
}
