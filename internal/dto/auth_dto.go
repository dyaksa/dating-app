package dto

type RegisterRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	BirthDate string `json:"birth_date" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
