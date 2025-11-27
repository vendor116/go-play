package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vendor116/playgo/pkg/openapi"
)

var _ ServiceInterface = (*service)(nil)

type service struct {
	infoHandler
}

func newService() ServiceInterface {
	return &service{
		infoHandler: infoHandler{},
	}
}

func SetupService(debug bool) *gin.Engine {
	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	openapi.RegisterHandlers(
		r.Group("/v1"),
		openapi.NewStrictHandler(newService(), nil),
	)

	return r
}
