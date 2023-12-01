package v1

import (
	"fmt"
	"net/http"

	"infra-server/config"
	"infra-server/routes/v1/service"
	"infra-server/routes/v1/upload"

	"github.com/gin-gonic/gin"
)

func Handle(route *gin.RouterGroup, cfg *config.ServerConfig, s service.Service) {
	route.POST("service", func(ctx *gin.Context) {
		body := &service.CreateServiceBody{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			fmt.Print(err)
			ctx.Status(400)
			return
		}

		created, err := s.CreateService(body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		// nginxConfig := s.CreateService(body)
		ctx.JSON(200, created)
	})

	route.GET("service", func(ctx *gin.Context) {
		vals, err := s.GetServices()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		ctx.JSON(200, vals)
	})

	route.GET("service/:id", func(ctx *gin.Context) {
		id, exists := ctx.Params.Get("id")
		if !exists {
			ctx.Status(400)
			return
		}
		vals, err := s.GetService(id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		ctx.JSON(200, vals)
	})

	route.POST("/service/:id/environment", func(ctx *gin.Context) {
		body := &service.CreateEnvironmentBody{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			fmt.Print(err)
			ctx.Status(400)
			return
		}

		id, exists := ctx.Params.Get("id")
		if !exists {
			ctx.Status(400)
			return
		}

		created, err := s.CreateEnvironment(id, body)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		}

		// nginxConfig := s.CreateService(body)
		ctx.JSON(200, created)
	})

	upload.Handle(route, cfg)
}
