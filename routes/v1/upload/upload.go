package upload

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"infra-server/config"

	"github.com/gin-gonic/gin"
)

type BindFile struct {
	Name string                `form:"name" binding:"required"`
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func Handle(router *gin.RouterGroup, cfg *config.ServerConfig) {
	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
			return
		}

		file := bindFile.File
		fileName := filepath.Base(file.Filename)
		fullPath := fmt.Sprintf("%s/%s", cfg.TempDir, fileName)

		if err := c.SaveUploadedFile(file, fullPath); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.Status(201)
	})
}
