package server

import (
	"github.com/lashkapashka/divideBill/internal/api"
)

type Server struct {
	api *api.API
}

func New() *Server {
	server := Server{
		api.New(),
	}

	return &server
}

func (s *Server) Run() {
	s.api.Run(":8085")
}