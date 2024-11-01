package interfaces

import (
	"context"

	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/entity"
	"github.com/google/uuid"
)

type PremiumPackagesRepository interface {
	ExistPackage(ctx context.Context, id uuid.UUID) (bool, error)
}

type UserPurchasesRepository interface {
	Purchases(ctx context.Context, entity *entity.UserPurchases) error
	IsPremium(ctx context.Context, userID uuid.UUID) (bool, error)
}
