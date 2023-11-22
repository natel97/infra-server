package cloudflare

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"infra-server/config"

	"gorm.io/gorm"
)

type ServiceConfig struct {
	IPAddress string
	Refresh   uint
}

type v1Service struct {
	repository Repository
	api        API
	config     *ServiceConfig
}

func getCloudflareV1Service() Service {
	return &v1Service{}
}

// func (service *v1Service) CreateRecord() error {
// 	return nil
// }

func (service *v1Service) Create(cfg *config.WebsiteConfig) error {
	domains, err := service.repository.GetZones()

	var match *Zone
	for _, domain := range domains {
		if strings.Contains(cfg.Domain, domain.Domain) {
			match = &domain
		}
	}

	if match == nil {
		return errors.New("No matching domain")
	}

	subdomain, err := service.repository.GetRecords(match.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if len(subdomain) > 0 {
		return errors.New("Domain exists")
	}

	err = service.api.CreateRecord(match.ID, dnsRecordAPI{
		RecordValue: service.config.IPAddress,
		Type:        "A",
		Proxied:     true,
		Name:        cfg.Domain,
		TTL:         1,
	})

	if err != nil {
		return err
	}

	service.repository.CreateRecord()

	return nil
}

func (service *v1Service) GetAll() ([]config.WebsiteConfig, error) {

	return nil, nil
}

// func (service *v1Service) UpdateRecord() error {
// 	return nil
// }

// func (service *v1Service) DeleteRecord() error {
// 	return nilServerConfig
// }

// func (service *v1Service) GetRecords() ([]DNSRecord, error) {
// 	records, err := service.repository.GetRecords()
// 	return records, err
// }

func (service *v1Service) refresh() error {
	fmt.Println("Refreshing domains")
	zones, err := service.api.GetZones()
	if err != nil {
		return err
	}

	for _, zone := range zones {
		records := service.api.GetRecords(zone.ID)
		fmt.Println(records)
	}

	return nil
}

func NewV1Service(repo Repository, api API, config *ServiceConfig) *v1Service {
	service := &v1Service{
		api:        api,
		repository: repo,
	}

	ticker := time.NewTicker(time.Duration(config.Refresh) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				err := service.refresh()
				if err != nil {
					fmt.Println("Error refreshing domains: ", err)
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return service
}
