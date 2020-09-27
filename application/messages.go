package application

import (
	"bufio"
	"os"
	"strings"
)

type Process struct {
	Id   string
	Ip   string
	Port string
}

type Message struct {
	S Process
	R string
	M string
}

func GetInfo(client Process) Message {
	//get the message from user and pack in into Message struct
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	t := strings.Fields(text)
	var m Message
	if t[0] == "EXIT" {
		m.R = "server"
		m.M = "EXIT"
		m.S = client
	}else{
		m.R = t[1]
		m.M = t[2]
		m.S = client
	}
	return m
}