package loadbalancer

import "natelubitz.com/config"

type Service interface {
	Create(config *config.LoadBalancer) error
	Update(config *config.LoadBalancer) error
	Delete(config *config.LoadBalancer) error
	GetAll() []config.LoadBalancer
	Get(id string) config.LoadBalancer
}
