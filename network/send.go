package network

import (
	"../application"
	"encoding/gob"
	"log"
	"net"
)

func UnicastSend(destination application.Process, m application.Message) {
	//Sends message to the destination process.
	c, err := net.Dial("tcp", destination.Ip+":"+destination.Port) //connect to TCP server
	if err != nil {
		log.Panic(err)
	}
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
	}
	_ = encoder.Encode(msg)
	return
}

