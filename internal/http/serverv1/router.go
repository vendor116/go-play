package serverv1

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vendor116/go-play/internal/http/middleware"
	"github.com/vendor116/go-play/pkg/openapi/goplayv1"
)

func RegisterRoutes(s goplayv1.StrictServerInterface) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	v1 := r.Group("/v1")
	v1.Use([]gin.HandlerFunc{
		middleware.SlogLogger(slog.Default()),
	}...)

	goplayv1.RegisterHandlers(v1, goplayv1.NewStrictHandler(s, nil))

	return r
}
