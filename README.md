# MP2

Simulate a simple chat room application that supports only private message

## To Run

### Send messages between different clients

Open three terminals. In the first terminal, start a TCP server and provide the port number.

```
go run server.go 8080
```

Then in the second terminal, start the first client, provide host address, port number, and username.

```
go run client.go 127.0.0.1 8080 user1
```

Then in the third terminal, do similar things:

```
go run client.go 127.0.0.1 8080 user2
```

Follow the instruction and send a message to another user.

```
send user1 hi
```

In user1's terminal, you will see:

```
Received 'hi' from user2
```

### Stop client or server (need to be implemented)

In the first terminal, run following command:

```
EXIT
```

Then the server will shut down, and each client will receives a notification.

## Structure and Design

### application

In `/application/message.go`, there are two struct:

```
type Process struct {
	Id   string
	Ip   string
	Port string
}

type Message struct {
	S Process
	R string
	M string
}
```

The `Process` struct contains the basic information of a process, and the `Message` struct contains four elements: `S` represent source, `R` represent receiver, `M` represent the text being sent.\

It mainly simulates the application layer. The `GetInfo` function takes a process as input, instruct the user to send the message, and then returns a `Message` struct.

### network

The `/network/connection.go` contains `Server` and `handleConnecion` function. Basically, The `Server` function starts a TCP listener, and then use `handleConnection` to handle multiple clients cuncurrently. `handleConnection` function can pass the `message` and `Conn` type to channels, so that they can be used to communicate in the main thread.

The`/network/send.go` contains `UnicastSend` function, which send the `message` through given `conn` connection with the support of `gob`.

### client and server

Initially, the client will send an empty `message` to the server, so that server can record its username and `conn` to a map. After that, the client program takes user inputs and send the message to the server. The server check whether it wants to exit or not. Then the server send the message to the destination user.\
For the server, we have a seperate thread that controls the ***EXIT*** property. We implement this property with the help of go channel and switch. If the user inputs `exit`, it will pass into the go channel and so `case <- control` section will be activated. Then it sends terminaion signal to all the connected clients by loop through the map we created previously.

## Deployment
* [Channels and Go Routines](https://www.justindfuller.com/2020/01/go-things-i-love-channels-and-goroutines/)
* [Create a TCP and UDP Client and Server using Go](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)
* [Go Routines](https://golangbot.com/goroutines/)
* [MP0 by Jiahong Li and Zheng Zhou](https://github.com/jiahongli18/DistributedSystemsMP0)
* [Net Package](https://golang.org/pkg/net/)
* [How to stop a Go Routine](https://stackoverflow.com/questions/6807590/how-to-stop-a-goroutine/6807784#6807784)


## Authors

* **Gary Ge** - *initial work*
