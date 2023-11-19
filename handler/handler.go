package handler

import (
	"natelubitz.com/config"
)

type Handler interface {
	Create(*config.LoadBalancer) error
	Update(*config.LoadBalancer) error
	Delete(*config.LoadBalancer) error
}
