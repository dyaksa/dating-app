package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/entity"
	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/repository"
	"github.com/google/uuid"
)

type purchasesUsecaseImpl struct {
	repository *repository.Repository
}

func NewPurchasesUsecase(repository *repository.Repository) PurchasesUsecase {
	return &purchasesUsecaseImpl{
		repository: repository,
	}
}

func (p *purchasesUsecaseImpl) Purchases(ctx context.Context, userID uuid.UUID, payload *dto.PurchasesRequest) error {
	existPackage, err := p.repository.PremiumPackage.ExistPackage(ctx, payload.PackageID)
	if err != nil {
		return err
	}

	if !existPackage {
		return fmt.Errorf("package with id %s not found", payload.PackageID)
	}

	return p.repository.UserPurchases.Purchases(ctx, &entity.UserPurchases{
		UserID:         userID,
		PackageID:      payload.PackageID,
		PurchaseDate:   time.Now(),
		ExpirationDate: time.Now().AddDate(0, 1, 0),
	})
}
