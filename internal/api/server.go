package api

import (
	"github.com/vendor116/playgo/internal/generated"
)

var _ generated.ServerInterface = (*Server)(nil)

type Server struct {
	*InfoHandler
}

func NewServer() *Server {
	return &Server{
		InfoHandler: NewInfoHandler(),
	}
}
