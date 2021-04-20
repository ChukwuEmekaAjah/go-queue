package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ChukwuEmekaAjah/go-queue/queue"
)

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

	q := new(queue.Queue)
	counter := 1

	go func() {
		for {
			n1 := strconv.Itoa(counter)
			q.Enqueue(n1)
			counter += 1
		}

	}()

	for {

		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go tcpHandler(c, q)
	}
}

func tcpHandler(c net.Conn, q *queue.Queue) {

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

		head := q.Dequeue()

		if head != (queue.Node{}) {
			c.Write([]byte(head.GetValue(head)))
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := "\n" + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
