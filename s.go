package main

import (
	"fmt"

	"github.com/ChukwuEmekaAjah/go-queue"
)

func main() {

	server := queue.Server{}

	go server.Create(":3000", "push")

	count := 1
	for {
		server.Send("hello world")
		count += 1
		fmt.Println("Count is", count)
	}

}
