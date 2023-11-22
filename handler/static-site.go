package handler

import "infra-server/config"

type staticSite struct {
}

func (staticSite *staticSite) Create(site *config.LoadBalancer) {
	// Upload zip
	// Unpack zip
	// Create NGINX Config
	// Create smlnk
	// Restart NGINX
}

func (staticSite *staticSite) Update(site *config.LoadBalancer) {
	// Delete previous directory
	// Upload zip
	// Unpack zip
}

func (staticSite *staticSite) Delete(site *config.LoadBalancer) {
	// Delete directory
	// Delete nginx config file
	// Restart NGINX
}
