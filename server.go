package MP2

import (
	"log"
	"net"
	"os"
	"./network"
	"./application"
)

func main() {
	//construct TCP channels for all
	//initialize the message channel to communicate
	var client application.Process
	arguments := os.Args
	//Starts the server
	client.Port = arguments[1]
	messages := make(chan application.Message)


	//go routine handle for each process

	//go channel to communicate and send message to destination
	for{
		go network.Server(client, messages)
		mes := <-messages
		// need to use mes.R to find destination.
		processMsg(mes.S, destination)
	}
}



func processMsg(from application.Process, to application.Process){
	var m application.Message
	m = application.GetInfo(from)
	network.UnicastSend(to, m)
}