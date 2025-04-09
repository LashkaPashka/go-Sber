package server

import (
	"log"

	"github.com/lashkapshka/go-Sber/internal/api"
	"github.com/lashkapshka/go-Sber/pkg/queue"
)

type Server struct {
	api *api.API
	kafka *queue.Client
}

func New() *Server {
	kaf, err := queue.New(
		[]string{"localhost:9092"},
		"topic-factors",
		"consumer-group",
	)

	if err != nil {
		log.Println("не получается иниициализировать kafka")
		return nil
	}

	server := Server{}

	server.api = api.New()
	server.kafka = kaf

	return &server
}

func (s *Server) Run() {
	s.api.Run(":8050")
}

func (s *Server) Consumer() {
	go s.kafka.Consumer()
}