package main

import (
	"fmt"
	"net"
	"os"
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

func (c *Client) Pull() {
	buffer := make([]byte, 8)
	for {

		_, err := c.connection.Read(buffer)

		if err != nil {
			panic(err)
		}
		fmt.Println("->: ", string(buffer))
	}
}

func (c *Client) Subscribe(topic string) {
	fmt.Fprintf(c.connection, topic)
}

func (c *Client) Disconnect() {
	fmt.Println("Closing connection to remote")
	c.connection.Close()
}

func main() {
	arguments := os.Args

	defaultServerAddress := "localhost:1996"

	if len(os.Args) > 1 {
		defaultServerAddress = arguments[1]
	}

	client := Client{nil}
	_, err := client.Connect(defaultServerAddress)

	if err != nil {
		fmt.Printf("Could not connect to remote server")
		panic(err)
	}

	client.Pull()

	// client.Disconnect()
}
