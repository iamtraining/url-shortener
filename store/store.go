package store

import (
	"fmt"

	"github.com/iamtraining/url-shortener/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	models.UrlStore
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		if err != nil {
			return nil, fmt.Errorf("failure while opening database: %w", err)
		}
		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("error while connecting to database: %w", err)
		}
	}

	return &Store{
		UrlStore: &UrlStore{DB: db},
	}, nil
}
