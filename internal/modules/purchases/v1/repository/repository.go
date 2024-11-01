package repository

import (
	"database/sql"

	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository/interfaces"
	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository/postgres"
)

type Repository struct {
	UserPurchases  interfaces.UserPurchasesRepository
	PremiumPackage interfaces.PremiumPackagesRepository
}

func NewPurchasesRepository(db *sql.DB) *Repository {
	return &Repository{
		UserPurchases:  postgres.NewUserPurchasesRepository(db),
		PremiumPackage: postgres.NewPremiumPackagesRepository(db),
	}
}
