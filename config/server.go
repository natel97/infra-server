package config

type ServerConfig struct {
	Port                  uint   `mapstructure:"PORT"`
	StaticSiteDirectory   string `mapstructure:"STATIC_SITE_DIRECTORY"`
	TempDir               string `mapstructure:"TEMP_DIR"`
	HTTPPort              uint   `mapstructure:"HTTP_PORT"`
	HTTPSPort             uint   `mapstructure:"HTTPS_PORT"`
	LoadBalancerDirectory string `mapstructure:"LOAD_BALANCER_DIRECTORY"`
	Environment           string `mapstructure:"ENVIRONMENT"`
	MaxUploadSize         uint   `mapstructure:"MAX_UPLOAD_SIZE_MB"`
}
