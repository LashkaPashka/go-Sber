package main

import (
	"log"

	"github.com/lashkapashka/divideBill/internal/server"
	"github.com/lashkapashka/divideBill/internal/service"
)

func main() {
	service := service.New()

	service.Divide()
}


func main1() {
	server := server.New()

	log.Println("Server's running 8085")
	server.Run()
}