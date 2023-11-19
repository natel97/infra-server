package service

import "natelubitz.com/config"

type service interface {
	Create(*config.WebsiteConfig) error
	Update(*config.WebsiteConfig) error
	Delete(*config.WebsiteConfig) error
}

type serviceHandler struct {
	services map[string]service
}

func (handler *serviceHandler) AddService(name string, service service) {
	handler.services[name] = service
}

func CreateServiceHandler() *serviceHandler {
	return &serviceHandler{
		services: map[string]service{},
	}
}

func CreateFromConfig(cfg *config.WebsiteConfig) {

}
