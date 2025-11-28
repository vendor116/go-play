package serverv1

import (
	"github.com/vendor116/go-play/internal/http/serverv1/info"
)

type Server struct {
	infoHandlers
}

func NewServer() *Server {
	return &Server{
		infoHandlers: info.NewHandlers(),
	}
}
