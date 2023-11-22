package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var apiZone = "https://api.cloudflare.com/client/v4/zones"
var apiBaseURL = "https://api.cloudflare.com/client/v4/zones/%s/dns_records"

type cloudflareError struct {
}

type cloudflareRespone[T interface{}] struct {
	Result  []T               `json:"result"`
	Success bool              `json:"success"`
	Errors  []cloudflareError `json:"errors"`
}

type zone struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Name   string `json:"name"`
}

type dnsRecordAPI struct {
	ID          string `json:"id"`
	RecordValue string `json:"content"`
	Name        string `json:"name"`
	Proxied     bool   `json:"proxied"`
	Type        string `json:"type"`
	Comment     string `json:"comment"`
	TTL         uint   `json:"ttl"`
}

func (api *v1API) GetZones() ([]zone, error) {
	bytes, err := api.handleAPICall("GET", apiZone, nil)
	if err != nil {
		return nil, err
	}

	response := cloudflareRespone[zone]{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

func (api *v1API) handleDNSAPICall(zone string, method string, route string, body []byte) ([]byte, error) {
	url := fmt.Sprintf(apiBaseURL, zone)
	if route != "" {
		url = fmt.Sprintf("%s/%s", url, route)
	}

	return api.handleAPICall(method, url, body)
}

func (api *v1API) handleAPICall(method string, route string, body []byte) ([]byte, error) {
	bodyReader := bytes.NewReader(body)

	req, err := http.NewRequest(method, route, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", api.config.Token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return []byte{}, readErr
	}

	return body, nil
}

//go:generate mockgen -source=api.go -destination=api_mock.go -package=cloudflare
type API interface {
	GetZones() ([]zone, error)
	CreateRecord(zone string, record dnsRecordAPI) error
	UpdateRecord(zone string, id string, record dnsRecordAPI) error
	DeleteRecord(zone string, id string) error
	GetRecords(zone string) error
	GetRecord(zone string, id string) error
}

type Config struct {
	Token string
}

type v1API struct {
	config *Config
}

func (api *v1API) CreateRecord(zone string, record dnsRecordAPI) error {
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	bytes, err := api.handleDNSAPICall(zone, "POST", "", body)
	if err != nil {
		return err
	}

	response := dnsRecordAPI{}
	json.Unmarshal(bytes, &response)

	return nil
}

func (api *v1API) UpdateRecord(zone string, id string, record dnsRecordAPI) error {
	body, err := json.Marshal(record)
	if err != nil {
		return err
	}

	bytes, err := api.handleDNSAPICall(zone, "PUT", "", body)

	if err != nil {
		return err
	}

	response := dnsRecordAPI{}
	json.Unmarshal(bytes, &response)

	return nil
}

func (api *v1API) DeleteRecord(zone string, id string) error {
	bytes, err := api.handleDNSAPICall(zone, "DELETE", id, nil)
	if err != nil {
		return err
	}

	response := dnsRecordAPI{}
	json.Unmarshal(bytes, &response)

	return nil
}

func (api *v1API) GetRecords(zone string) error {
	bytes, err := api.handleDNSAPICall(zone, "GET", "", nil)
	if err != nil {
		return err
	}

	response := dnsRecordAPI{}
	json.Unmarshal(bytes, &response)

	return nil
}

func (api *v1API) GetRecord(zone string, id string) error {
	bytes, err := api.handleDNSAPICall(zone, "GET", "", nil)
	if err != nil {
		return err
	}

	response := dnsRecordAPI{}
	json.Unmarshal(bytes, &response)

	return nil
}

func NewAPI(config *Config) *v1API {
	return &v1API{
		config: config,
	}
}
