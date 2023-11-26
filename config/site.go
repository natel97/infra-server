package config

type WebsiteConfig struct {
	Domain   string `json:"domain"`
	Comments string `json:"comments"`
	Name     string `json:"name"`

	LoadBalancer LoadBalancer `json:"loadBalancer"`
}

type LoadBalancer struct {
	Type         string       `json:"type"`
	StaticConfig StaticConfig `json:"staticConfig"`
	ProxyConfig  ProxyConfig  `json:"proxyConfig"`
}

type Environment struct {
	Name string `json:"name"`
}

type StaticConfig struct {
	IsSPA bool `json:"isSPA"`
}

type ProxyConfig struct {
	ProxyFrom string `json:"proxyFrom"`
}
