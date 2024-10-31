package config

import (
	"database/sql"

	"github.com/dyaksa/dating-app/internal/infra/database/postgres"
)

func NewPostgres() *sql.DB {
	return postgres.Connect()
}
