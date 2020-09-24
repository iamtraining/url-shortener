package store

import (
	"fmt"

	"github.com/iamtraining/url-shortener/models"
	"github.com/jmoiron/sqlx"
)

type UrlStore struct {
	*sqlx.DB
}

func (s *UrlStore) Create(u *models.Url) error {
	if err := s.Get(u, `INSERT INTO url_shortener (id, original ,short) VALUES ($1,$2,$3) RETURNING *`,
		u.ID, u.Original, u.Short); err != nil {
		return fmt.Errorf("failure while creating url entity %w", err)
	}
	return nil
}

func (s *UrlStore) Read(id string) (models.Url, error) {
	var url models.Url
	if err := s.Get(&url, `SELECT * FROM url_shortener WHERE id = $1`, id); err != nil {
		return url, fmt.Errorf("failure while getting url entity %w", err)
	}
	return url, nil
}

func (s *UrlStore) Update(u *models.Url) error {
	if err := s.Get(u, `UPDATE url_shortener SET original = $1, short = $2 WHERE id = $3 RETURNING *`,
		u.Original, u.Short, u.ID); err != nil {
		return fmt.Errorf("failure while updating url entity %w", err)
	}
	return nil
}

func (s *UrlStore) Delete(id string) error {
	if _, err := s.Exec(`DELETE FROM url_shortener WHERE id = $1`, id); err != nil {
		return fmt.Errorf("error while deleting url entity %w", err)
	}
	return nil
}
