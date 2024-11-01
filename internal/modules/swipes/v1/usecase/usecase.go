package usecase

import (
	"context"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/google/uuid"
)

type SwipeUsecase interface {
	GetSwipeUser(ctx context.Context, userID uuid.UUID) (*dto.SwipesUserResponse, error)
	SwipeUser(ctx context.Context, swiperID uuid.UUID, payload *dto.SwipesRequest) error
}
