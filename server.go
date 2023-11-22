package main

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"infra-server/config"
	v1 "infra-server/routes/v1"
	"infra-server/routes/v1/service"
	"infra-server/services/dns/cloudflare"
	"infra-server/services/load-balancer/nginx"
	"infra-server/utils"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func readPipe(reader io.Reader, prefix string) {
	r := bufio.NewReader(reader)
	var outStr string
	var line []byte
	for {
		line, _, _ = r.ReadLine()
		if line != nil {
			outStr = string(line)
			fmt.Println(prefix + outStr)
		}
	}
}

//go:embed frontend/dist/*
var frontend embed.FS

func EmbedReact(urlPrefix, buildDirectory string, em embed.FS) gin.HandlerFunc {
	dir := static.LocalFile(buildDirectory, true)
	embedDir, _ := fs.Sub(em, buildDirectory)
	fileserver := http.FileServer(http.FS(embedDir))

	if urlPrefix != "" {
		fileserver = http.StripPrefix(urlPrefix, fileserver)
	}

	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "/api/") {
			c.Status(404)
			return
		}
		if !dir.Exists(urlPrefix, c.Request.URL.Path) {
			c.Request.URL.Path = "/"
		}
		fileserver.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	}
}

func attachFrontend(server *gin.Engine, dev bool) {
	if !dev {
		server.Use(EmbedReact("/", "frontend/dist", frontend))
	} else {
		command := exec.Command("npm", "run", "dev")
		command.Dir = "./frontend"

		cmdOut, _ := command.StdoutPipe()
		cmdErr, _ := command.StderrPipe()

		go readPipe(cmdOut, "\033[0;35m[Frontend] \033[0m")
		go readPipe(cmdErr, "\033[0;31m[Frontend] \033[0m")
		err := command.Start()

		if err != nil {
			fmt.Println(err)
		}
	}
}

func attachAPIs(group *gin.RouterGroup, cfg *config.ServerConfig, service service.Service) {
	v1Group := group.Group("/v1")
	v1.Handle(v1Group, cfg, service)
}

func loggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()
			log["level"] = "access"

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func loadENV() *config.ServerConfig {
	cfg := &config.ServerConfig{}
	prodConfig := viper.New()
	prodConfig.SetConfigName("config")
	prodConfig.SetConfigType("env")
	prodConfig.AddConfigPath(".")
	err := prodConfig.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	localConfig := viper.New()
	localConfig.SetConfigName("local")
	localConfig.SetConfigType("env")
	localConfig.AddConfigPath(".")
	localConfig.ReadInConfig()

	prodConfig.MergeConfigMap(localConfig.AllSettings())
	prodConfig.Unmarshal(&cfg)

	return cfg
}

type migratable interface {
	Migrate(db *gorm.DB) error
}

func migrate(db *gorm.DB, repositories []migratable) error {
	for _, repository := range repositories {
		err := repository.Migrate(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(loggerMiddleware())
	cfg := loadENV()

	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	cloudflareRepository := cloudflare.NewV1Repository(db)

	migrate(db, []migratable{
		cloudflareRepository,
	})

	api := cloudflare.NewAPI(&cloudflare.Config{
		Token: cfg.CloudflareAPIToken,
	})

	ip := utils.GetPublicIP()

	cloudflareService := cloudflare.NewV1Service(cloudflareRepository, api, &cloudflare.ServiceConfig{
		IPAddress: ip,
		Refresh:   10,
	})

	nginxService := nginx.NewNginxHandler(cfg)
	service := service.NewV1Service(cloudflareService, nginxService)

	apiGroup := server.Group("api")
	server.MaxMultipartMemory = int64(cfg.MaxUploadSize << 20)
	attachAPIs(apiGroup, cfg, service)

	attachFrontend(server, cfg.Environment == "development")
	fmt.Println("Server started on port: ", cfg.Port)
	server.Run(fmt.Sprintf(":%d", cfg.Port))
}
