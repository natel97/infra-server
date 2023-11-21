package cloudflare

import (
	"fmt"
	"time"
)

type v1Service struct {
	repository Repository
	api        API
}

func getCloudflareV1Service() Service {
	return &v1Service{}
}

func (service *v1Service) CreateRecord() error {
	return nil
}

func (service *v1Service) UpdateRecord() error {
	return nil
}

func (service *v1Service) DeleteRecord() error {
	return nil
}

func (service *v1Service) GetRecords() ([]DNSRecord, error) {
	records, err := service.repository.GetRecords()
	return records, err
}

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

func NewV1Service(repo Repository, api API, refresh uint) *v1Service {
	service := &v1Service{
		api:        api,
		repository: repo,
	}

	ticker := time.NewTicker(time.Duration(refresh) * time.Second)
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
