package main

import (
	"./application"
	"./network"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
)

func dial(server application.Process)(c net.Conn){
	address := server.Ip + ":" + server.Port
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	return c
}

func receive(c net.Conn, control chan bool){
	for {
		decoder := gob.NewDecoder(c)
		//Decode the message from server
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		if mes.M == "EXIT" {
			fmt.Println("Exit the program.")
			c.Close()
			control <- true
			break
		}
		fmt.Printf("Received '%s' from %s\n", mes.M, mes.S.Id)
	}
}

func main(){
	arguments := os.Args
	var server application.Process
	var self application.Process
	server.Ip = arguments[1]
	server.Port = arguments[2]
	self.Id = arguments[3]
	conn := dial(server)
	control := make(chan bool)
	initialMessage := application.Message{S: self, R: "server"}
	network.UnicastSend(conn, initialMessage)
	go receive(conn, control)
	fmt.Print("please send application in this pattern: send 'username' 'YourMessage'\n")
	println("enter EXIT to quit.")
	for{
		select{
		case <- control:
			//termination signal from server, close as well.
			conn.Close()
			return
		default:
			m:= application.GetInfo(self)
			network.UnicastSend(conn, m)
			if m.M =="EXIT"{
				println("connection closed.")
				return
			}
		}

	}
}