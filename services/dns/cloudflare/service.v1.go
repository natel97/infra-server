package cloudflare

type cloudflareV1 struct {
}

func getCloudflareV1Service() Service {
	return &cloudflareV1{}
}

func (service *cloudflareV1) CreateRecord() error { return nil }
func (service *cloudflareV1) UpdateRecord() error { return nil }
func (service *cloudflareV1) DeleteRecord() error { return nil }
func (service *cloudflareV1) GetRecords()         {}
