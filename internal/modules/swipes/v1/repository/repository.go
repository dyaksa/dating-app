package repository

import (
	"database/sql"

	purchases "github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository/interfaces"
	purchasesRepo "github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository/postgres"
	swipes "github.com/dyaksa/dating-app/internal/modules/swipes/v1/repository/interfaces"
	swipesRepo "github.com/dyaksa/dating-app/internal/modules/swipes/v1/repository/postgres"
	user "github.com/dyaksa/dating-app/internal/modules/user/v1/repository/interfaces"
	userRepo "github.com/dyaksa/dating-app/internal/modules/user/v1/repository/postgres"
)

type Repository struct {
	Swipes    swipes.SwipesRepository
	User      user.UserRepository
	Purchases purchases.UserPurchasesRepository
}

func NewSwipesRepository(db *sql.DB) *Repository {
	return &Repository{
		Swipes:    swipesRepo.NewSwipeRepository(db),
		User:      userRepo.NewUserRepo(db),
		Purchases: purchasesRepo.NewUserPurchasesRepository(db),
	}
}
