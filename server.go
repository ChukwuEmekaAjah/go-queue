package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"github.com/chukwuemekaajah/go-queue/queue"
)

fmt.Println(queue.Node{time.Now(), "hello", strconv.Itoa(234), nil})
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

func main() {
	arguments := os.Args

	portAddress := ":1996"

	if len(arguments) > 1 {
		portAddress = ":" + arguments[1]
	}

	l, err := net.Listen("tcp", portAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	queue := new(Queue)
	counter := 1

	go func() {
		for {
			n1 := Node{time.Now(), "hello", strconv.Itoa(counter), nil}
			queue.Enqueue(n1)
			counter += 1
		}

	}()

	for {

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go tcpHandler(c, queue)
	}
}

func tcpHandler(c net.Conn, queue *Queue) {

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		head := queue.Dequeue()

		if head != (Node{}) {
			c.Write([]byte(head.value))
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := "\n" + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
