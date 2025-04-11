package server

import (
	"github.com/lashkapashka/divideBill/internal/api"

	//"github.com/lashkapashka/divideBill/pkg/queue"
)

type Server struct {
	api *api.API
	//rabbit *queue.RabbitMQ
}

func New() *Server {
	server := Server{
		api: api.New(),
		//rabbit: queue.New(),
	}

	return &server
}

func (s *Server) Run() {
	s.api.Run(":8085")
}

// func (s *Server) Consumer() {
// 	s.rabbit.Consumer()
// }