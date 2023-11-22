package filegen

import (
	"fmt"

	"infra-server/config"
)

var staticSiteConfig string = `
server {
  server_name %s;
  access_log /var/log/nginx/%s.access.log json_combined;

  root %s;
  location / {
    try_files $uri $uri/ %s =404;
  }

  listen 443;
}

server {
  if ($host = %s) {
    return 301 https://$host$request_uri;
  }

  listen 80;
  return 404;
}
`

type nginxStaticSite struct {
	config *config.ServerConfig
}

func (server *nginxStaticSite) GenerateFile(site *config.WebsiteConfig) string {
	fullPath := fmt.Sprintf("%s/%s", server.config.StaticSiteDirectory, site.Name)
	indexPath := fmt.Sprintf("%s/index.html", fullPath)
	if !site.LoadBalancer.StaticConfig.IsSPA {
		indexPath = ""
	}

	nginxString := fmt.Sprintf(staticSiteConfig, site.Domain, site.Name, fullPath, indexPath, site.Domain)
	return nginxString
}

func NewStaticSiteGenerator(serverConfig *config.ServerConfig) *nginxStaticSite {
	return &nginxStaticSite{
		config: serverConfig,
	}
}
