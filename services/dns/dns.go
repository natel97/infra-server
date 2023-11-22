package dns

import "infra-server/config"

type Repository interface{}

type Service interface {
	Create(*config.WebsiteConfig) error
	GetAll() ([]config.WebsiteConfig, error)
}
