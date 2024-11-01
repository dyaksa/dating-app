package dto

type UpdateProfileRequest struct {
	Bio               string `json:"bio" binding:"required"`
	Location          string `json:"location" binding:"required"`
	ProfilePictureURL string `json:"profile_picture_url" binding:"required"`
}
