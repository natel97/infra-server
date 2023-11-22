package cloudflare

type Zone struct {
	ID     string
	Domain string
}

type DNSRecord struct {
	ID     string
	Zone   string
	Domain string
}

type Repository interface {
	GetZones() ([]Zone, error)
	GetRecords(string) ([]DNSRecord, error)
	GetAllRecords() ([]DNSRecord, error)
	UpdateRecord()
	DeleteZone()
	DeleteRecord()
	CreateZone()
	CreateRecord()
}
