package service

import (
	"database/sql"
)

type Service struct {
	Db *sql.DB
}

func New() *Service {
	return &Service{}
}
