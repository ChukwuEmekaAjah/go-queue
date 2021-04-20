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

type Queue struct {
	head  *Node
	tail  *Node
	count int
}

func (c *Queue) Enqueue(item Node) int {

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

// func main() {

// 	queue := new(Queue)
// 	n1 := Node{time.Now(), "hello", "meat", nil}
// 	n2 := Node{time.Now(), "hello", "miss", nil}
// 	n3 := Node{time.Now(), "hello", "world", nil}

// 	queue.Enqueue(n1)
// 	queue.Enqueue(n2)
// 	queue.Enqueue(n3)
// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.Dequeue())
// 	fmt.Println(queue.Dequeue())
// }
