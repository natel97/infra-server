package service

import (
	"fmt"

	"infra-server/config"
	"infra-server/services/dns"
	loadbalancer "infra-server/services/load-balancer"
)

type Service interface {
	CreateService(cfg *config.WebsiteConfig) error
	GetServices() ([]config.WebsiteConfig, error)
}

type v1Service struct {
	// services map[string]service
	dnsService          dns.Service
	loadBalancerService loadbalancer.Service
}

func NewV1Service(dnsService dns.Service, loadBalancerService loadbalancer.Service) *v1Service {
	return &v1Service{
		dnsService:          dnsService,
		loadBalancerService: loadBalancerService,
	}
}

func (service *v1Service) CreateService(cfg *config.WebsiteConfig) error {
	err := service.loadBalancerService.Create(cfg)
	if err != nil {
		return err
	}

	err = service.dnsService.Create(cfg)
	if err != nil {
		return err
	}

	return nil
}

func (service *v1Service) GetServices() ([]config.WebsiteConfig, error) {
	balancer, err := service.loadBalancerService.GetAll()
	if err != nil {
		return nil, err
	}

	dns, err := service.dnsService.GetAll()
	if err != nil {
		return nil, err
	}

	fmt.Println(balancer, dns)
	return nil, nil
}
