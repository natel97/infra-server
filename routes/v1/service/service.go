package service

import (
	"errors"
	"infra-server/routes/v1/upload"
	"infra-server/services/dns"
	loadbalancer "infra-server/services/load-balancer"
	"mime/multipart"

	"github.com/google/uuid"
)

type Service interface {
	CreateService(cfg *CreateServiceBody) (*GetServiceResponse, error)
	GetServices() ([]GetServiceResponse, error)
	GetService(id string) (*GetSingleServiceResponse, error)
	CreateEnvironment(id string, environment *CreateEnvironmentBody) (*EnvironmentStub, error)
	UploadAndPromoteVersion(deploymentID string, environmentID string, file *multipart.FileHeader) error
	SetCurrentVersion(deploymentID string, environmentID string, deployID string) error
	CreateURL(serviceID string, domain CreateURLBody) error
}

func (service *v1Service) CreateURL(serviceID string, domain CreateURLBody) error {
	// TODO: do
	return nil
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
			Name:  "Port",
			Type:  "number",
			Value: 3000,
		},
		{
			ID:    "2",
			Name:  "Healthcheck Endpoint",
			Type:  "string",
			Value: "/health",
		},
		{
			ID:    "3",
			Name:  "Docker Image",
			Type:  "string",
			Value: "",
		},
	},
}

type v1Service struct {
	// services map[string]service
	dnsService          dns.Service
	fileHandler         upload.FileManager
	loadBalancerService loadbalancer.Service
	inMemoryServices    []GetSingleServiceResponse
}

func NewV1Service(dnsService dns.Service, loadBalancerService loadbalancer.Service, fileHandler upload.FileManager) *v1Service {
	return &v1Service{
		dnsService:          dnsService,
		loadBalancerService: loadBalancerService,
		fileHandler:         fileHandler,
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
		ID:           uuid.NewString(),
		Name:         details.Name,
		Environments: []EnvironmentStub{},
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

func (service *v1Service) GetService(id string) (*GetSingleServiceResponse, error) {
	for _, value := range service.inMemoryServices {
		if value.ID == id {
			return &value, nil
		}
	}

	return nil, nil
}

func (service *v1Service) CreateEnvironment(id string, environment *CreateEnvironmentBody) (*EnvironmentStub, error) {
	var match *GetSingleServiceResponse
	index := 0
	for i, s := range service.inMemoryServices {
		if s.ID == id {
			match = &s
			index = i
		}
	}

	if match == nil {
		return nil, nil
	}

	env := EnvironmentStub{
		ID:     uuid.NewString(),
		Name:   environment.Name,
		Status: "new",
	}

	match.Environments = append(match.Environments, env)
	service.inMemoryServices[index] = *match
	return &env, nil
}

func (service *v1Service) UploadAndPromoteVersion(deploymentID string, environmentID string, file *multipart.FileHeader) error {
	version, err := service.fileHandler.SaveNewVersion(deploymentID, environmentID, file)
	if err != nil {
		return err
	}

	for deploymentIndex, deploy := range service.inMemoryServices {
		if deploy.ID != deploymentID {
			continue
		}
		deploy.AvailableDeployments = append(deploy.AvailableDeployments, AvailableDeployment{
			ID:        version.ID,
			Timestamp: *version.DateUploaded,
		})

		service.inMemoryServices[deploymentIndex] = deploy
	}
	err = service.SetCurrentVersion(deploymentID, environmentID, version.ID)

	if err != nil {
		return err
	}

	return nil
}

func (service *v1Service) SetCurrentVersion(deploymentID string, environmentID string, deployID string) error {
	err := service.fileHandler.PromoteVersion(deploymentID, environmentID, deployID)
	if err != nil {
		return err
	}

	return nil
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
