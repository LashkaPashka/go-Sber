package main

import (
	"fmt"
	"log"

	"github.com/lashkapshka/go-Sber/internal/server"
	"github.com/lashkapshka/go-Sber/internal/service"
)

func main() {
	service := service.New()

	data := service.DivideBill("name1")

	fmt.Println(data)
}


func main1() {
	server := server.New()
	
	//Kafka
	log.Println("Kafka's running")	
	server.Consumer()

	// Server's running
	log.Println("Server's running")
	server.Run()


}