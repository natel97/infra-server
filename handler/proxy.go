package handler

import "infra-server/config"

type proxySite struct {
}

func (proxySite *proxySite) Create(site *config.LoadBalancer) {
	// Upload zip
	// Unpack zip
	// Create load balancer
	// Create smlnk
	// Restart load balancer
}

func (proxySite *proxySite) Update(site *config.LoadBalancer) {
	// Delete previous directory
	// Upload zip
	// Unpack zip
}

func (proxySite *proxySite) Delete(site *config.LoadBalancer) {
	// Delete directory
	// Delete nginx config file
	// Restart NGINX
}
