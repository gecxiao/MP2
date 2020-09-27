package network

import (
	"../application"
	"encoding/gob"
	"net"
)

func UnicastSend(c net.Conn, m application.Message) {
	//Sends message to the destination process.
	encoder := gob.NewEncoder(c)
	msg := application.Message{
		S: m.S,
		R: m.R,
		M: m.M,
	}
	_ = encoder.Encode(msg)
	return
}