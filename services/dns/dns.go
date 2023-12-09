package dns

import "infra-server/config"

type Repository interface{}

type Domain struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Service interface {
	CreateSubdomain(*config.WebsiteConfig) error
	GetSubdomains(domainID string) ([]config.WebsiteConfig, error)
	GetDomains() []Domain
	GetDomain(id string) (Domain, error)
}
