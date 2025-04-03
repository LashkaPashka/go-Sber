package main

import (
	"log"

	"github.com/lashkapshka/go-Sber/internal/server"
)

func main() {
	server := server.New()
	
	//Kafka
	log.Println("Kafka's running")	
	server.Consumer()

	// Server's running
	log.Println("Server's running")
	server.Run()


}