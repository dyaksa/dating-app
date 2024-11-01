package interfaces

import (
	"context"
	"time"

	"github.com/dyaksa/dating-app/internal/modules/auth/v1/domain"
)

type TokenRepository interface {
	Generate(ctx context.Context, payload *domain.TokenClaim, exipred time.Duration) (string, error)
}
