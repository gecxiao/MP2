package network

import (
	"../application"
	"encoding/gob"
	"fmt"
	"net"
)

func Transfer(c net.Conn, m application.Message){
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
	}
	_ = encoder.Encode(msg)
}

func handleConnection(c net.Conn, messages chan application.Message, conns chan map[string]net.Conn) {
	temp := make(map[string]net.Conn)
	for {
		decoder := gob.NewDecoder(c)
		mes := new(application.Message)
		_ = decoder.Decode(mes)
		println(mes.M)
		println(mes.S.Id)
		temp[mes.S.Id] = c
		conns <- temp
		messages <- *mes
		if mes.M == "EXIT" {
			c.Close()
			break
		}
		//transfer(c, *mes)
	}
}

func Server(server application.Process, messages chan application.Message, conns chan map[string]net.Conn) {
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
