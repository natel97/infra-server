package cloudflare

import (
	"gorm.io/gorm"
)

type cloudflareZone struct {
	gorm.Model
	ID     string
	Status string
	Name   string
}

type cloudflareDNSRecord struct {
	gorm.Model
	ID          string
	RecordValue string
	Name        string
	Proxied     bool
	Type        string
	Comment     string
	TTL         uint
	ZoneID      string
}

type v1Repository struct {
	db *gorm.DB
}

func (repo *v1Repository) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&cloudflareZone{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&cloudflareDNSRecord{})
	if err != nil {
		return err
	}

	return nil
}

func NewV1Repository(db *gorm.DB) *v1Repository {
	return &v1Repository{
		db: db,
	}
}

func (repo *v1Repository) GetZones() ([]Zone, error) {
	dbZones := []cloudflareZone{}
	tx := repo.db.Find(&dbZones)

	if tx.Error != nil {
		return nil, tx.Error
	}

	zones := []Zone{}
	for _, z := range dbZones {
		zones = append(zones, Zone{
			ID:     z.ID,
			Domain: z.Name,
		})
	}

	return zones, nil
}

func (repo *v1Repository) CreateZone(zone Zone) error {
	dbZone := &cloudflareZone{
		ID:   zone.ID,
		Name: zone.Domain,
	}
	err := repo.db.FirstOrCreate(&dbZone, "id = ?", zone.ID).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo *v1Repository) GetRecords(zone string) ([]DNSRecord, error) {
	dbRecords := []cloudflareDNSRecord{}
	tx := repo.db.Where("id = ?", zone).Scan(&dbRecords)

	if tx.Error != nil {
		return nil, tx.Error
	}

	zones := []DNSRecord{}
	for _, r := range dbRecords {
		zones = append(zones, DNSRecord{
			ID:     r.ID,
			Domain: r.Name,
			Zone:   r.ZoneID,
		})
	}

	return zones, nil
}

func (repo *v1Repository) GetAllRecords() ([]DNSRecord, error) {
	dbRecords := []cloudflareDNSRecord{}
	tx := repo.db.Find(&dbRecords)

	if tx.Error != nil {
		return nil, tx.Error
	}

	zones := []DNSRecord{}
	for _, r := range dbRecords {
		zones = append(zones, DNSRecord{
			ID:     r.ID,
			Domain: r.Name,
			Zone:   r.ZoneID,
		})
	}

	return zones, nil
}

func (repo *v1Repository) UpdateRecord() {

}

func (repo *v1Repository) DeleteZone() {

}

func (repo *v1Repository) DeleteRecord() {

}

func (repo *v1Repository) CreateRecord() {

}
