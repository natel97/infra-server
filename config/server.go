package config

type ServerConfig struct {
	Port                uint   `mapstructure:"PORT"`
	StaticSiteDirectory string `mapstructure:"STATIC_SITE_DIRECTORY"`
	TempDir             string `mapstructure:"TEMP_DIR"`
	HTTPPort            uint   `mapstructure:"HTTP_PORT"`
	HTTPSPort           uint   `mapstructure:"HTTPS_PORT"`

	NginxSitesAvailable string `mapstructure:"NGINX_SITES_AVAILABLE"`
	NginxSitesEnabled   string `mapstructure:"NGINX_SITES_ENABLED"`
	NginxRestartCommand string `mapstructure:"NGINX_RESTART_COMMAND"`

	Environment        string `mapstructure:"ENVIRONMENT"`
	MaxUploadSize      uint   `mapstructure:"MAX_UPLOAD_SIZE_MB"`
	CloudflareAPIToken string `mapstructure:"CLOUDFLARE_API_TOKEN"`
}
