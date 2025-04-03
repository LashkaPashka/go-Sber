package main

import (
	"fmt"

	"github.com/lashkapshka/go-Sber/pkg/queue"
)

func main() {
	cli, _ := queue.New(
		[]string{"localhost:9092"},
		"test-topic",
		"consumer-group",
	)
	
	cli.Producer("Client: balasananrafael")
	fmt.Println("сообщение отправлено в очередь")
}