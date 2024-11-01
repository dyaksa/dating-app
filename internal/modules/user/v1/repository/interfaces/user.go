package interfaces

import (
	"context"

	"github.com/dyaksa/dating-app/internal/modules/user/v1/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(ctx context.Context, entity *entity.User) error
	ExistUser(ctx context.Context, email string) (*entity.User, error)
	SwipeUser(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}
