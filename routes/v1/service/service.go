package service

import (
	"errors"
	"infra-server/services/dns"
	loadbalancer "infra-server/services/load-balancer"

	"github.com/google/uuid"
)

type Service interface {
	CreateService(cfg *CreateServiceBody) (*GetServiceResponse, error)
	GetServices() ([]GetServiceResponse, error)
}

var SettingsForType = map[string][]DeploySetting{
	"static-website": {
		{
			ID:    "1",
			Name:  "Is SPA",
			Type:  "boolean",
			Value: false,
		},
	},
	"kubernetes-deployment": {
		{
			ID:    "1",
			Name:  "Is SPA",
			Type:  "boolean",
			Value: false,
		},
	},
}

type v1Service struct {
	// services map[string]service
	dnsService          dns.Service
	loadBalancerService loadbalancer.Service
	inMemoryServices    []GetSingleServiceResponse
}

func NewV1Service(dnsService dns.Service, loadBalancerService loadbalancer.Service) *v1Service {
	return &v1Service{
		dnsService:          dnsService,
		loadBalancerService: loadBalancerService,
		inMemoryServices:    []GetSingleServiceResponse{},
	}
}

func (service *v1Service) CreateService(details *CreateServiceBody) (*GetServiceResponse, error) {
	// err := service.loadBalancerService.Create(cfg)
	// if err != nil {
	// 	return err
	// }

	// err = service.dnsService.Create(cfg)
	// if err != nil {
	// 	return err
	// }

	// return nil
	matchingSettings, exists := SettingsForType[details.Type]
	if !exists {
		return nil, errors.New("type does not exist")
	}

	newService := GetSingleServiceResponse{
		ID:   uuid.NewString(),
		Name: details.Name,
		DeploymentSettings: DeploySettings{
			ID:       uuid.NewString(),
			Type:     details.Type,
			Settings: matchingSettings,
		},
	}
	service.inMemoryServices = append(service.inMemoryServices, newService)

	return &GetServiceResponse{
		ID:           newService.ID,
		Name:         newService.Name,
		Type:         newService.DeploymentSettings.Type,
		Visibility:   "Public", // TODO: temporary
		Environments: []EnvironmentStub{},
	}, nil
}

func (service *v1Service) GetServices() ([]GetServiceResponse, error) {
	// balancer, err := service.loadBalancerService.GetAll()
	// if err != nil {
	// 	return nil, err
	// }

	// cfg := []config.WebsiteConfig{}
	// for _, s := range balancer {
	// 	cfg = append(cfg, config.WebsiteConfig{
	// 		LoadBalancer: s,
	// 	})
	// }

	// return cfg, nil
	response := []GetServiceResponse{}
	for _, item := range service.inMemoryServices {
		response = append(response, GetServiceResponse{
			ID:           item.ID,
			Name:         item.Name,
			Type:         item.DeploymentSettings.Type,
			Visibility:   "Public", // TODO: Temporary
			Environments: item.Environments,
		})
	}

	return response, nil
}
