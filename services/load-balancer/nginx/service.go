package nginx

import (
	"errors"
	"fmt"

	"infra-server/config"
	"infra-server/services/adapters"
	"infra-server/services/load-balancer/nginx/filegen"
)

type FileGen interface {
	GenerateFile(*config.WebsiteConfig) string
}

type nginx struct {
	config            *config.ServerConfig
	loaders           map[string]FileGen
	fileSystemHandler adapters.FileSystemHandler
}

func NewNginxHandler(config *config.ServerConfig) *nginx {
	balancer := nginx{}
	balancer.loaders = map[string]FileGen{}
	balancer.loaders["static-website"] = filegen.NewStaticSiteGenerator(config)
	balancer.loaders["proxy"] = filegen.NewProxyGenerator(config)
	// TODO: enable mock add to deps
	balancer.fileSystemHandler = adapters.NewRealFileSystemHandler()
	balancer.config = config
	return &balancer
}

func (n *nginx) Create(config *config.WebsiteConfig) error {
	loader := n.loaders[config.LoadBalancer.Type]
	if loader == nil {
		return errors.New("config type is invalid")
	}

	file := loader.GenerateFile(config)
	fileName := fmt.Sprintf("%s/%s.conf", n.config.NginxSitesAvailable, config.Domain)
	err := n.fileSystemHandler.CreateFile(fileName, file)
	if err != nil {
		return err
	}

	enabledPath := n.config.NginxSitesEnabled

	if n.config.NginxSitesEnabled[0] == "."[0] {
		cwd, err := n.fileSystemHandler.GetCWD()
		if err != nil {
			return err
		}
		enabledPath = fmt.Sprintf("%s%s", cwd, enabledPath[1:])
	}

	linkFile := fmt.Sprintf("%s/%s.conf", enabledPath, config.Domain)

	err = n.fileSystemHandler.CopyFile(fileName, linkFile)
	if err != nil {
		return err
	}

	err = n.fileSystemHandler.RunCommand(n.config.NginxRestartCommand)
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

func (n *nginx) GetAll() ([]config.LoadBalancer, error) { return nil, nil }

func (n *nginx) Get(id string) *config.LoadBalancer {
	return nil
}
