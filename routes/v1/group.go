package v1

import (
	"fmt"
	"net/http"

	"infra-server/config"
	"infra-server/routes/v1/service"
	"infra-server/routes/v1/upload"

	"github.com/gin-gonic/gin"
)

const KB = uint(1024)
const MB = 1024 * KB

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

	route.POST("/service/:id/url", func(ctx *gin.Context) {
		var body service.CreateURLBody

		serviceID, exists := ctx.Params.Get("id")

		if !exists {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if err := ctx.ShouldBind(&body); err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		err := s.CreateURL(serviceID, body)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	})

	route.POST("/upload", func(c *gin.Context) {
		var bindFile upload.BindFile

		if c.Request.ContentLength > int64(cfg.MaxUploadSize*MB) {
			c.Status(http.StatusRequestEntityTooLarge)
			return
		}

		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		file := bindFile.File
		err := s.UploadAndPromoteVersion(bindFile.Service, bindFile.Environment, file)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.Status(201)
	})
}
