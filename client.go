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
	decoder := gob.NewDecoder(c)
	//Decode the message from server
	mes := new(application.Message)
	_ = decoder.Decode(mes)
	fmt.Printf("Received '%s' from %s", mes.M, mes.S.Id)
	for {
		if mes.M == "EXIT" {
			fmt.Println("Exit the program.")
			return
		}
		fmt.Println(mes.M)
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
	go receive(conn)
	for{
		m:= application.GetInfo(self)
		if m.M == "EXIT"{
			return
		}
		//termination signal from server, close as well.
		network.UnicastSend(conn, m)
	}
}