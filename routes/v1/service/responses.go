package service

import (
	"time"
)

type CreateURLBody struct {
	EnvironmentID string `json:"environmentId"`
	Name          string `json:"name"`
	DomainID      string `json:"domainId"`
	Subdomain     string `json:"subdomain"`
}

type EnvironmentStub struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	LastDeploy string `json:"lastDeploy"`
}

type GetServiceResponse struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         string            `json:"type"`
	Visibility   string            `json:"visibility"`
	Environments []EnvironmentStub `json:"environments"`
}

type EnvironmentDeployment struct {
	ID          string `json:"id"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

type DeployLog struct {
	AvailableDeploymentID string
	Environments          []string `json:"environments"`
}

type AvailableDeployment struct {
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type DeploySetting struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type DeploySettings struct {
	ID       string          `json:"id"`
	Type     string          `json:"type"`
	Settings []DeploySetting `json:"settings"`
}

type GetSingleServiceResponse struct {
	ID                   string                `json:"id"`
	Name                 string                `json:"name"`
	AvailableDeployments []AvailableDeployment `json:"availableDeployments"`
	Environments         []EnvironmentStub     `json:"environments"`
	DeploymentSettings   DeploySettings        `json:"deploymentSettings"`
	DeployHistory        []DeployLog           `json:"deployHistory"`
	EnvironmentVariables []EnvironmentVariable `json:"environmentVariables"`
	VaultEntries         []VaultRecord         `json:"vaultEntries"`
	Services             []Provision           `json:"services"`
	URLs                 []URL                 `json:"urls"`
}

type EnvironmentVariable struct {
	ID           string   `json:"id"`
	Key          string   `json:"key"`
	Value        string   `json:"value"`
	Environments []string `json:"environments"`
	Active       []string `json:"active"`
}

type VaultRecord struct {
	ID           string   `json:"id"`
	Environments []string `json:"environments"`
	Active       []string `json:"active"`
	Length       uint     `json:"length"`
	VisibleHint  string   `json:"visibleHint"`
}

type Provision struct {
	ID               string   `json:"id"`
	ServiceID        string   `json:"serviceID"`
	ServiceName      string   `json:"serviceName"`
	ExposedVariables []string `json:"exposedVariables"`
}

type URL struct {
	ID              string  `json:"id"`
	EnvironmentName string  `json:"environmentName"`
	Visibility      string  `json:"visibility"`
	Link            DNSLink `json:"link"`
}

type DNSLink struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	Type       string `json:"type"`
	RedirectTo string `json:"redirectTo"`
}

type CreateEnvironmentBody struct {
	Name string `json:"name"`
}

type CreateServiceBody struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
