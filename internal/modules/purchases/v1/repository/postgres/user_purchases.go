package postgres

import (
	"context"
	"database/sql"

	"github.com/dyaksa/dating-app/internal/modules/purchases/v1/entity"
	"github.com/google/uuid"
)

type UserPurchasesRepository struct {
	db *sql.DB
}

func NewUserPurchasesRepository(db *sql.DB) *UserPurchasesRepository {
	return &UserPurchasesRepository{
		db: db,
	}
}

func (r *UserPurchasesRepository) Purchases(ctx context.Context, entity *entity.UserPurchases) error {
	query := `INSERT INTO user_purchases (user_id, package_id, purchase_date, expiration_date) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, entity.UserID, entity.PackageID, entity.PurchaseDate, entity.ExpirationDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserPurchasesRepository) IsPremium(ctx context.Context, userID uuid.UUID) (bool, error) {
	query := `SELECT pp.swipe_limit FROM user_purchases up 
	INNER JOIN premium_packages pp ON up.package_id = pp.id
	WHERE up.user_id = $1 AND up.expiration_date > NOW()`

	var swipeLimit bool
	err := r.db.QueryRowContext(ctx, query, userID).Scan(&swipeLimit)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return swipeLimit, nil
}
