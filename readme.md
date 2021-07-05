<h1 align="center">Go-queue</h1>

A naive Go implementation of a Kafka queue. It enables you to create `pub/sub` or `push/pull` queues. It uses a TCP connection to enable communication between clients and servers. You can create data publishers and data consumers using any of the above two patterns.


## What you get
- The ability to create a message queue that stores data in memory 
- It only sends data to available connections and when there are no consumers, it keeps growing the message list in-memory.
- Create multiple consumers listening to publishes using any of the `push/pull` or `pub/sub` methods.


## How to use
```bash
go get github.com/ChukwuEmekaAjah/go-queue
```

```go

```

## API 
Methods on the message publisher. It has only two methods
- publisher.Create(portAddress string, socketType string)
- publisher.Send(data string)

Methods on the message consumer. It has only three methods
- consumer.Connect(serverAddress string)
- consumer.Pull()
- consumer.Subscribe(topic string)



## Contributing
In case you have any ideas, features you would like to be included or any bug fixes, you can send a PR.

- Clone the repo

```bash
git clone https://github.com/ChukwuEmekaAjah/go-queue.git
```

## Todo
- Improve on documentation with examples and use cases
