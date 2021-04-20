package queue

import (
	"time"
)

type Node struct {
	dateAdded time.Time
	topic     string
	value     string
	next      *Node
}

func (n *Node) GetValue(node Node) string {
	return node.value
}

type Queue struct {
	head  *Node
	tail  *Node
	count int
}

func (c *Queue) Enqueue(value string) int {

	item := Node{time.Now(), "hello", value, nil}
	if c.count == 0 {
		c.head = &item
		c.tail = &item
	} else {
		c.tail.next = &item

		c.tail = &item
	}

	c.count += 1
	return c.count
}

func (c *Queue) Dequeue() Node {

	head := c.head

	c.head = c.head.next

	return *head
}
