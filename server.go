package main

import (
	"./application"
	"./network"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	//construct TCP channels for all
	//initialize the message channel to communicate
	var client application.Process
	arguments := os.Args
	//Starts the server
	client.Port = arguments[1]
	control :=make(chan bool)
	conns := make(chan map[string]net.Conn)
	messages := make(chan application.Message)
	//go channel to communicate and send message to destination]
	go handleExit(client, control, messages, conns)
	for{
		reader := bufio.NewReader(os.Stdin)
		var cmd string
		cmd, _ = reader.ReadString('\n')
		if strings.TrimSpace(cmd) == "EXIT" {
			fmt.Println("Server is exiting...")
			//Sends the termination signal to all the connected clients
			control <- true
			time.Sleep(10)
			return
		}
	}
}

func handleExit(server application.Process, control chan bool, messages chan application.Message,
	conns chan map[string]net.Conn) {
	for {
		select {
		case <- control:
			return
		default:
			go network.Server(server, messages, conns)
			mes := <- messages
			cm := <- conns
			network.Transfer(cm[mes.R], mes)
		}
	}
}
