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
	control :=make(chan bool, 1)
	conns := make(chan net.Conn, 1)
	messages := make(chan application.Message, 1)
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
	conns chan net.Conn) {
	cm := make(map[string]net.Conn)
	go network.Server(server, messages, conns)
	for {
		select {
		case <- control:
			shutdown := new(application.Message)
			shutdown.M = "server is closed."
			for username, conn := range cm{
				shutdown.R = username
				network.UnicastSend(conn, *shutdown)
				conn.Close()
			}
			return
		default:
			mes := <- messages
			if mes.R == "server"{
				c := <- conns
				cm[mes.S.Id] = c
			}else {
				if cm[mes.R] != nil{
					network.UnicastSend(cm[mes.R], mes)
				}else{
					errorMessage := application.Message{
						M: "The user you want to send is not connected",
					}
					network.UnicastSend(cm[mes.S.Id], errorMessage)
				}
			}
		}
	}
}
