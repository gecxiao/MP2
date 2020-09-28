package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"net"
)

func handleConnection(c net.Conn, cm map[string]net.Conn, server application.Process) {
	for {
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		if mes.M == "EXIT"{
			println(mes.S.Id+" is disconnected.")
			cm[mes.S.Id] = nil
		}else{
			Transfer(cm, server, *mes)
		}
	}
}

func ExitUsers(cm map[string]net.Conn, control chan bool){
	shutdown := new(application.Message)
	shutdown.M = "EXIT"
	for username, conn := range cm{
		shutdown.R = username
		UnicastSend(conn, *shutdown)
	}
	control <- true
	return
}

func Transfer(cm map[string]net.Conn, server application.Process, mes application.Message) {
	if cm[mes.R] != nil {
		UnicastSend(cm[mes.R], mes)
	} else {
		errorMessage := application.Message{
			M: "The user you want to send is not connected",
			S: server,
		}
		UnicastSend(cm[mes.S.Id], errorMessage)
	}
}

func Server(server application.Process, cm map[string]net.Conn, control chan bool) {
	//input: the network# and the # of connections it will receive
	//listen to the client and decode the application, then send via channel
	//simulate the delay here.
	ln, err := net.Listen("tcp", server.Ip+":"+server.Port) //creates server
	if err != nil {
		fmt.Println(err)
	}
	defer ln.Close()
	for {

		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		cm[mes.S.Id] = c
		//handle concurrently
		go handleConnection(c, cm, server)
		<- control
		ln.Close()
	}
}
