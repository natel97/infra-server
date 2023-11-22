package loadbalancer

import "infra-server/config"

type Service interface {
	// Create(config *config.WebsiteConfig) error
	// Update(config *config.LoadBalancer) error
	// Delete(config *config.WebsiteConfig) error
	// GetAll() ([]config.LoadBalancer, error)
	// Get(id string) *config.LoadBalancer

	Create(config *config.WebsiteConfig) error
	Update(config *config.WebsiteConfig) error
	Delete(config *config.WebsiteConfig) error
	GetAll() ([]config.LoadBalancer, error)
	Get(id string) *config.LoadBalancer
}
