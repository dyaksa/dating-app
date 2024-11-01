package usecase

import (
	"context"

	"github.com/dyaksa/dating-app/internal/dto"
)

type Authusecase interface {
	Register(ctx context.Context, payload *dto.RegisterRequest) error
	Login(ctx context.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error)
}
