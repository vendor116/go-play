package serverv1

import (
	"context"

	"github.com/vendor116/go-play/pkg/openapi/goplayv1"
)

type infoHandlers interface {
	GetInfo(context.Context, goplayv1.GetInfoRequestObject) (goplayv1.GetInfoResponseObject, error)
}
