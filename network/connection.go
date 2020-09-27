package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"net"
)

func handleConnection(c net.Conn, messages chan application.Message, conns chan net.Conn) {
	for {
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		if mes.R == "server"{
			if mes.M == "EXIT" {
				messages <- *mes
				conns <- nil
				break
			}
			conns <- c
		} else{
			messages <- *mes
		}
	}
	c.Close()
}

func Server(server application.Process, messages chan application.Message, conns chan net.Conn) {
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
		go handleConnection(c, messages, conns)
	}
}
