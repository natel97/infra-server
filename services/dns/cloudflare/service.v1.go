package cloudflare

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"infra-server/config"
	"infra-server/services/dns"

	"gorm.io/gorm"
)

type ServiceConfig struct {
	IPAddress      string
	RefreshSeconds uint
}

type v1Service struct {
	repository Repository
	api        API
	config     *ServiceConfig
}

func (service *v1Service) CreateSubdomain(cfg *config.WebsiteConfig) error {
	domains, err := service.repository.GetZones()

	if err != nil {
		return err
	}

	var match *Zone
	for _, domain := range domains {
		if strings.Contains(cfg.Domain, domain.Domain) {
			match = &domain
		}
	}

	if match == nil {
		return errors.New("no matching domain")
	}

	subdomain, err := service.repository.GetRecords(match.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if len(subdomain) > 0 {
		return errors.New("domain exists")
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

func (service *v1Service) GetSubdomains(id string) ([]config.WebsiteConfig, error) {
	records, err := service.repository.GetRecords(id)

	if err != nil {
		return nil, err
	}

	sites := []config.WebsiteConfig{}

	for _, record := range records {
		sites = append(sites, config.WebsiteConfig{
			Domain:       record.Domain,
			LoadBalancer: config.LoadBalancer{},
		})
	}
	return sites, nil
}

func (service *v1Service) GetDomains() []dns.Domain {
	zones, err := service.repository.GetZones()
	if err != nil {
		return []dns.Domain{}
	}

	domains := []dns.Domain{}
	for _, zone := range zones {
		domains = append(domains, dns.Domain{
			ID:  zone.ID,
			URL: zone.Domain,
		})
	}

	return domains
}

func (service *v1Service) GetDomain(id string) (dns.Domain, error) {
	zones, err := service.repository.GetZones()
	if err != nil {
		return dns.Domain{}, err
	}

	for _, zone := range zones {
		if zone.ID == id {
			return dns.Domain{
				ID:  zone.ID,
				URL: zone.Domain,
			}, nil
		}
	}

	// todo temporary logic, create repo query
	return dns.Domain{}, errors.New("not found")
}

func (service *v1Service) refresh() error {
	fmt.Println("Refreshing domains")
	zones, err := service.api.GetZones()
	fmt.Println(zones)
	if err != nil {
		return err
	}

	for _, zone := range zones {
		err := service.repository.CreateZone(Zone{
			ID:     zone.ID,
			Domain: zone.Name,
		})

		if err != nil {
			return err
		}
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

	ticker := time.NewTicker(time.Duration(config.RefreshSeconds) * time.Second)
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
