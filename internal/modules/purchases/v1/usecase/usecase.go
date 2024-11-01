package usecase

import (
	"context"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/google/uuid"
)

type PurchasesUsecase interface {
	Purchases(ctx context.Context, userID uuid.UUID, payload *dto.PurchasesRequest) error
}
