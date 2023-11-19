package loadbalancer

import "natelubitz.com/config"

type LoadBalancer interface {
	Create(config *config.LoadBalancer) error
	Update(config *config.LoadBalancer) error
	Delete(config *config.LoadBalancer) error
}
