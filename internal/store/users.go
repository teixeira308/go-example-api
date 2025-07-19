package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	DB *sql.DB
}

func (s *UsersStore) Create(ctx context.Context) error {
	return nil
}
