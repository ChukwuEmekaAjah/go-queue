package queue

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	connection net.Conn
}

func (c *Client) Connect(serverAddress string) (connection net.Conn, err error) {

	connection, err = net.Dial("tcp", serverAddress)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c.connection = connection
	return connection, nil
}

func (c *Client) Pull(handler func(data string)) {
	reader := bufio.NewReader(c.connection)
	for {

		data, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading from publisher")
			panic(err)
		}
		go handler(data[0 : len(data)-1])
		fmt.Println("->: ", data[0:len(data)-1])
	}
}

func (c *Client) Subscribe(topic string, handler func(data string)) {
	reader := bufio.NewReader(c.connection)
	for {

		data, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading from publisher")
			panic(err)
		}
		go handler(data[0 : len(data)-1])
		fmt.Println("->: ", data[0:len(data)-1])
	}
}

func (c *Client) Disconnect() {
	fmt.Println("Closing connection to remote")
	c.connection.Close()
}
