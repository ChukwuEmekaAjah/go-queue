package main

import (
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/ChukwuEmekaAjah/go-queue/queue"
)

type Server struct {
	listener    net.Listener
	socketType  string
	connections []net.Conn
	queue       queue.Queue
	receiver    chan string
}

func (s *Server) Create(portAddress string, socketType string) {

	listener, err := net.Listen("tcp", portAddress)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer listener.Close()

	s.listener = listener
	s.socketType = socketType
	s.queue = queue.Queue{}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		s.connections = append(s.connections, connection)
	}
}

func (s *Server) Send(data string) {
	s.queue.Enqueue(data)

	if len(s.connections) > 0 {
		head := s.queue.Dequeue()
		if s.socketType == "push" {
			s.connections[0].Write([]byte(head.GetValue()))
		} else if s.socketType == "pub" {
			for _, connection := range s.connections {
				connection.Write([]byte(head.GetValue()))
			}
		}
	}

	if len(s.connections) > 1 {
		s.connections = append(s.connections[1:], s.connections[0])
	}
}

// func (s *Server) handler(connection net.Conn) {
// 	fmt.Println("I am running now")
// 	for {
// 		head := s.queue.Dequeue()
// 		connection.Write([]byte(head.GetValue() + " "))
// 	}
// }

// has to have its internal queue

func main() {
	arguments := os.Args

	portAddress := ":1996"

	if len(arguments) > 1 {
		portAddress = ":" + arguments[1]
	}

	server := new(Server)

	go func() {
		counter := 1
		for {
			server.Send(strconv.Itoa(counter))
			counter += 1
		}
	}()

	server.Create(portAddress, "push")

	// l, err := net.Listen("tcp", portAddress)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer l.Close()

	// q := new(queue.Queue)
	// counter := 1

	// go func() {
	// 	for {
	// 		n1 := strconv.Itoa(counter)
	// 		q.Enqueue(n1)
	// 		counter += 1
	// 	}

	// }()

	// for {

	// 	c, err := l.Accept()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	go tcpHandler(c, q)
	// }
}

// func tcpHandler(c net.Conn, q *queue.Queue) {

// 	for {
// 		netData, err := bufio.NewReader(c).ReadString('\n')
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		if strings.TrimSpace(string(netData)) == "STOP" {
// 			fmt.Println("Exiting TCP server!")
// 			return
// 		}

// 		head := q.Dequeue()

// 		if head != (queue.Node{}) {
// 			c.Write([]byte(head.GetValue()))
// 		}

// 		fmt.Print("-> ", string(netData))
// 		t := time.Now()
// 		myTime := "\n" + t.Format(time.RFC3339) + "\n"
// 		c.Write([]byte(myTime))
// 	}
// }
