package MP2

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"./network"
	"./application"
	"os"
)

func main(){
	arguments := os.Args
	var server application.Process
	var self application.Process
	server.Ip = arguments[1]
	server.Port = arguments[2]
	self.Id = arguments[3]
	go dial(server)
	for{
		m:= application.GetInfo(self)
		if m.M == "EXIT"{
			break
		}
		//termination signal from server, close as well.
		network.UnicastSend(server, m)
	}
}

func dial(server application.Process){
	address := server.Ip + ":" + server.Port
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Panic(err)
	}
	decoder := gob.NewDecoder(c)
	//Decode the message from server
	mes := new(application.Message)
	_ = decoder.Decode(mes)

	for {
		if mes.M == "EXIT" {
			fmt.Println("Exit the program.")
			return
		}
		fmt.Println(mes.M)
	}
}