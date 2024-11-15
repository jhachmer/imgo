package store

import "database/sql"

type Store interface {
	Save() error
}

type DBStore struct {
	db sql.DB
}
