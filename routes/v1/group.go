package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"natelubitz.com/config"
	"natelubitz.com/routes/v1/service"
	"natelubitz.com/routes/v1/upload"
	"natelubitz.com/services/load-balancer/nginx"
)

func Handle(route *gin.RouterGroup, cfg *config.ServerConfig) {
	nginxService := nginx.NewNginxHandler(cfg)

	s := service.CreateServiceHandler()
	s.AddService("nginx", nginxService)
	route.POST("services", func(ctx *gin.Context) {
		body := &config.WebsiteConfig{}
		err := ctx.ShouldBindJSON(body)
		if err != nil {
			fmt.Print(err)
			ctx.Status(400)
			return
		}

		// nginxConfig := s.CreateService(body)
		ctx.JSON(200, map[string]interface{}{
			"config": body,
		})
	})

	upload.Handle(route, cfg)
}
