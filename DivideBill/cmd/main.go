package main

import (
	"log"

	"github.com/lashkapashka/divideBill/internal/server"
)

func main() {
	server := server.New()
	
	//log.Println("RabbitMQ's running")
	//server.Consumer()

	log.Println("Server's running 8085")
	server.Run()


}