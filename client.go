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

func receive(c net.Conn){
	for {
		decoder := gob.NewDecoder(c)
		//Decode the message from server
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		fmt.Printf("Received '%s' from %s\n", mes.M, mes.S.Id)
		if mes.M == "EXIT" {
			fmt.Println("Exit the program.")
			return
		}
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
	initialMessage := application.Message{S: self, R: "server"}
	network.UnicastSend(conn, initialMessage)
	go receive(conn)
	fmt.Print("please send application in this pattern: send 'username' 'YourMessage'\n")
	println("enter EXIT to quit.")
	for{
		m:= application.GetInfo(self)
		network.UnicastSend(conn, m)
		if m.M == "EXIT"{
			println("connection closed.")
			return
		}
		//termination signal from server, close as well.
	}
}