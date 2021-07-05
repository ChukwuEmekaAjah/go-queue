package queue

import (
	"fmt"
	"net"
)

type Server struct {
	listener    net.Listener
	socketType  string
	connections []net.Conn
	queue       Queue
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
	s.queue = Queue{}

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		s.connections = append(s.connections, connection)
	}
}

func (s *Server) Send(data string) {
	data += "\n"
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

	if len(s.connections) > 1 && s.socketType == "push" { // naive rotation of the connections
		s.connections = append(s.connections[1:], s.connections[0])
	}
}
