package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	Name         string    `db:"name"`
	Gender       string    `db:"gender"`
	BirthDate    string    `db:"birth_date"`
	IsVerified   bool      `db:"is_verified"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
