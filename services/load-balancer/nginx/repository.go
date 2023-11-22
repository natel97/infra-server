package nginx

import "gorm.io/gorm"

type Repository interface{}

type v1Repository struct {
	db *gorm.DB
}

type nginxEntry struct {
}

func (repo *v1Repository) Migrate() error {
	err := repo.db.AutoMigrate(&nginxEntry{})
	if err != nil {
		return err
	}

	return nil
}

func (repo *v1Repository) CreateRecord() error {
	// repo.db.Create()
	return nil
}

func NewV1Repository(db *gorm.DB) *v1Repository {
	return &v1Repository{
		db: db,
	}
}
