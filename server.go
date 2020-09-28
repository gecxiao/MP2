package main

import (
	"./application"
	"./network"
	"bufio"
	"net"
	"os"
	"strings"
)

func main() {
	//construct TCP channels for all
	//initialize the message channel to communicate
	var client application.Process
	var server application.Process
	cm := make(map[string]net.Conn)
	server.Id = "server"
	arguments := os.Args
	//Starts the server
	client.Port = arguments[1]
	server.Port = arguments[1]
	userControl :=make(chan bool, 1)
	lnControl :=make(chan bool, 1)
	//go channel to communicate and send message to destination]
	go getExit(userControl)
	go network.Server(server, cm, lnControl)
	<- userControl
	network.ExitUsers(cm,lnControl)
	return
}

func getExit(control chan bool){
	for{
		reader := bufio.NewReader(os.Stdin)
		var cmd string
		cmd, _ = reader.ReadString('\n')
		if strings.TrimSpace(cmd) == "EXIT" {
			//Sends the termination signal to all the connected clients
			control <- true
		}
	}
}



