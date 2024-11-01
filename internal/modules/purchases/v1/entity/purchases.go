package entity

import (
	"time"

	"github.com/google/uuid"
)

type PremiumPackages struct {
	ID                uuid.UUID `db:"id"`
	Name              string    `db:"name"`
	Price             int       `db:"price"`
	Description       string    `db:"description"`
	SwipeLimit        bool      `db:"swipe_limit"`
	VerificationLabel bool      `db:"verification_label"`
}

type UserPurchases struct {
	ID             uuid.UUID `db:"id"`
	UserID         uuid.UUID `db:"user_id"`
	PackageID      uuid.UUID `db:"package_id"`
	PurchaseDate   time.Time `db:"purchase_date"`
	ExpirationDate time.Time `db:"expiration_date"`
}
