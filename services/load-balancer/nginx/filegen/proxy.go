package filegen

import (
	"fmt"

	"natelubitz.com/config"
)

var proxySiteConfig string = `
server {
  server_name %s;
  client_max_body_size 100M;

  location / {
    proxy_pass %s;
    # add_header 'Access-Control-Allow-Origin' '*';
  }

  listen 443;
}

server {
  server_name %s;

  if ($host = %s) {
    return 301 https://$host$request_uri;
  }

  listen 80;
  return 404;
}
`

type nginxProxy struct {
	config *config.ServerConfig
}

func (server *nginxProxy) GenerateFile(site *config.WebsiteConfig) string {
	nginxString := fmt.Sprintf(proxySiteConfig, site.Domain, site.LoadBalancer.ProxyConfig.ProxyFrom, site.Domain, site.Domain)
	return nginxString
}

func NewProxyGenerator(serverConfig *config.ServerConfig) *nginxProxy {
	return &nginxProxy{
		config: serverConfig,
	}
}
