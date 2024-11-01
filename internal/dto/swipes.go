package dto

import "github.com/google/uuid"

type SwipesRequest struct {
	TargetProfileID uuid.UUID `json:"target_profile_id" binding:"required"`
	SwipeType       bool      `json:"swipe_type"`
}

type SwipesUserResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	BirthDate  string    `json:"birthdate"`
	IsVerified bool      `json:"is_verified"`
}
