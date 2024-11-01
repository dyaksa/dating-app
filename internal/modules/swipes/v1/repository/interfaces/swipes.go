package interfaces

import (
	"context"

	"github.com/dyaksa/dating-app/internal/modules/swipes/v1/entity"
	"github.com/google/uuid"
)

type SwipesRepository interface {
	Swipe(ctx context.Context, entity *entity.Swipes) error
	Count(ctx context.Context, userID uuid.UUID) (int, error)
}
