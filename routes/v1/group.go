package v1

import (
	"fmt"
	"net/http"

	"infra-server/config"
	"infra-server/routes/v1/service"
	"infra-server/routes/v1/upload"

	"github.com/gin-gonic/gin"
)

func Handle(route *gin.RouterGroup, cfg *config.ServerConfig, service service.Service) {
	route.POST("service", func(ctx *gin.Context) {
		body := &config.WebsiteConfig{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			fmt.Print(err)
			ctx.Status(400)
			return
		}

		err = service.CreateService(body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		// nginxConfig := s.CreateService(body)
		ctx.JSON(200, map[string]interface{}{
			"config": body,
		})
	})

	route.GET("service", func(ctx *gin.Context) {
		vals, err := service.GetServices()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}

		ctx.JSON(200, vals)
	})

	upload.Handle(route, cfg)
}
