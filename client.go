package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args

	defaultServerAddress := "localhost:1996"

	if len(os.Args) > 1 {
		defaultServerAddress = arguments[1]
	}

	c, err := net.Dial("tcp", defaultServerAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		text := "data" // keep reading data from the queue without stopping
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
