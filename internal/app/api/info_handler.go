package api

import (
	"context"

	"github.com/vendor116/playgo/pkg/openapi"
)

type infoHandler struct{}

func (infoHandler) GetInfo(_ context.Context, _ openapi.GetInfoRequestObject) (openapi.GetInfoResponseObject, error) {
	return openapi.GetInfo200JSONResponse{
		Name:    "play-go",
		Version: "dev",
	}, nil
}
