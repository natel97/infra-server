package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"natelubitz.com/config"
)

type service struct {
	db *sql.DB
}

func InitDB() (*service, error) {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		return nil, err
	}

	return &service{
		db: db,
	}, nil
}

func (s *service) CreateSite(cfg *config.WebsiteConfig) {

}
