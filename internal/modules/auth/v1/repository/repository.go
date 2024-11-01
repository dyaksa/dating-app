package repository

import (
	"database/sql"

	"github.com/dyaksa/dating-app/internal/modules/auth/v1/repository/interfaces"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/repository/jwt"
	user "github.com/dyaksa/dating-app/internal/modules/user/v1/repository/interfaces"
	"github.com/dyaksa/dating-app/internal/modules/user/v1/repository/postgres"
)

type Repository struct {
	User  user.UserRepository
	Token interfaces.TokenRepository
}

func NewAuthRepository(db *sql.DB) *Repository {
	return &Repository{
		User:  postgres.NewUserRepo(db),
		Token: jwt.NewJWTRepository(),
	}
}
