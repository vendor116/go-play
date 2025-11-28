package info

import (
	"context"

	"github.com/vendor116/go-play/pkg/openapi/goplayv1"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) GetInfo(_ context.Context, _ goplayv1.GetInfoRequestObject) (goplayv1.GetInfoResponseObject, error) {
	return goplayv1.GetInfo200JSONResponse{
		Name:    "play-go",
		Version: "dev",
	}, nil
}
