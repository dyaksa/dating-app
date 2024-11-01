package repository

import (
	"database/sql"

	user "github.com/dyaksa/dating-app/internal/modules/user/v1/repository/interfaces"
	userRepo "github.com/dyaksa/dating-app/internal/modules/user/v1/repository/postgres"
)

type Repository struct {
	User user.UserRepository
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		User: userRepo.NewUserRepo(db),
	}
}
