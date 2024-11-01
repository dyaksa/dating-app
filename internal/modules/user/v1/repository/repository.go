package repository

import (
	"database/sql"

	"github.com/dyaksa/dating-app/internal/modules/user/v1/repository/interfaces"
	"github.com/dyaksa/dating-app/internal/modules/user/v1/repository/postgres"
)

type Repository struct {
	User interfaces.UserRepository
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		User: postgres.NewUserRepo(db),
	}
}
