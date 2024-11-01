package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/domain"
	"github.com/dyaksa/dating-app/internal/modules/auth/v1/repository"
	"github.com/dyaksa/dating-app/internal/modules/user/v1/entity"
	"github.com/dyaksa/dating-app/internal/utils/password"
)

type authUsecaseImpl struct {
	repository *repository.Repository
}

func NewAuthUsecaseImpl(repository *repository.Repository) Authusecase {
	return &authUsecaseImpl{repository: repository}
}

func (uc *authUsecaseImpl) Register(ctx context.Context, payload *dto.RegisterRequest) error {
	existUser, err := uc.repository.User.ExistUser(ctx, payload.Email)
	if err != nil {
		return err
	}

	if existUser != nil {
		return errors.New("email already exist")
	}

	if err := uc.repository.User.Save(ctx, &entity.User{
		Name:         payload.Name,
		Email:        payload.Email,
		PasswordHash: payload.Password,
		Gender:       payload.Gender,
		BirthDate:    payload.BirthDate,
		IsVerified:   false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (uc *authUsecaseImpl) Login(ctx context.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := uc.repository.User.ExistUser(ctx, payload.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if ok := password.Verify(payload.Password, user.PasswordHash); !ok {
		return nil, errors.New("password not match")
	}

	tokenClaims := new(domain.TokenClaim)
	tokenClaims.Alg = "HS256"
	tokenClaims.User.ID = user.ID.String()
	token, err := uc.repository.Token.Generate(ctx, tokenClaims, time.Duration(1)*time.Hour)
	if err != nil {
		return nil, err
	}

	response := &dto.LoginResponse{
		Token: token,
		Email: user.Email,
	}

	return response, nil
}
