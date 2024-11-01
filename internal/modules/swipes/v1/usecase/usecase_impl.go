package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/dyaksa/dating-app/internal/dto"
	"github.com/dyaksa/dating-app/internal/modules/swipes/v1/entity"
	"github.com/dyaksa/dating-app/internal/modules/swipes/v1/repository"
	"github.com/google/uuid"
)

type swipeUsecaseImpl struct {
	repository *repository.Repository
}

func NewSwipeUsecase(repository *repository.Repository) SwipeUsecase {
	return &swipeUsecaseImpl{
		repository: repository,
	}
}

func (u *swipeUsecaseImpl) SwipeUser(ctx context.Context, swiperID uuid.UUID, payload *dto.SwipesRequest) error {

	count, err := u.repository.Swipes.Count(ctx, swiperID)
	if err != nil {
		return err
	}

	ok, err := u.repository.Purchases.IsPremium(ctx, swiperID)
	if err != nil {
		return err
	}

	if !ok && count >= 3 {
		return fmt.Errorf("swipe limit exceeded")
	}

	err = u.repository.Swipes.Swipe(ctx, &entity.Swipes{
		SwiperID:        swiperID,
		TargetProfileID: payload.TargetProfileID,
		SwipeType:       payload.SwipeType,
		CreatedAt:       time.Now(),
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *swipeUsecaseImpl) GetSwipeUser(ctx context.Context, userID uuid.UUID) (*dto.SwipesUserResponse, error) {
	existUser, err := u.repository.User.SwipeUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	if existUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	users := &dto.SwipesUserResponse{
		ID:         existUser.ID,
		Name:       existUser.Name,
		BirthDate:  existUser.BirthDate,
		Gender:     existUser.Gender,
		IsVerified: existUser.IsVerified,
	}
	return users, nil
}
