package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"net"
)
func handleConnection(c net.Conn, messages chan application.Message){
	for{
		decoder := gob.NewDecoder(c)

		mes := new(application.Message)
		_ = decoder.Decode(mes)
		if mes.M == "EXIT" {
			break
		}
		messages <- *mes
	}
}
func Server(server application.Process, messages chan application.Message) {
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
		//handle concurrently
		go handleConnection(c, messages)
	}
}