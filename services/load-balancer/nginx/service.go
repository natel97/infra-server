package nginx

import (
	"errors"
	"fmt"
	"os"

	"natelubitz.com/config"
	"natelubitz.com/services/load-balancer/nginx/filegen"
)

type FileGen interface {
	GenerateFile(*config.WebsiteConfig) string
}

type fsHandler interface {
	CreateFile(name string, content string) error
}

type nginx struct {
	config            *config.ServerConfig
	loaders           map[string]FileGen
	fileSystemHandler fsHandler
}

type realFSHandler struct{}

func (handler *realFSHandler) CreateFile(name string, content string) error {
	return os.WriteFile(name, []byte(content), 0644)
}

func NewNginxHandler(config *config.ServerConfig) *nginx {
	balancer := nginx{}
	balancer.loaders = map[string]FileGen{}
	balancer.loaders["static-website"] = filegen.NewStaticSiteGenerator(config)
	balancer.loaders["proxy"] = filegen.NewProxyGenerator(config)
	balancer.fileSystemHandler = &realFSHandler{}
	return &balancer
}

func (n *nginx) Create(config *config.WebsiteConfig) error {
	loader := n.loaders[config.LoadBalancer.Type]
	if loader == nil {
		return errors.New("config type is invalid")
	}

	file := loader.GenerateFile(config)
	fileName := fmt.Sprintf("%s/%s.conf", n.config.LoadBalancerDirectory, config.Domain)
	err := n.fileSystemHandler.CreateFile(fileName, file)
	if err != nil {
		return err
	}

	return nil
}

func (n *nginx) Update(config *config.WebsiteConfig) error {
	return nil
}

func (n *nginx) Delete(config *config.WebsiteConfig) error {
	return nil
}

func (n *nginx) GetAll() []config.LoadBalancer { return nil }

func (n *nginx) Get(id string) *config.LoadBalancer {
	return nil
}
