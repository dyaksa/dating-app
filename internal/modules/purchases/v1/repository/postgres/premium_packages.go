package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type PremiumPackagesRepository struct {
	db *sql.DB
}

func NewPremiumPackagesRepository(db *sql.DB) *PremiumPackagesRepository {
	return &PremiumPackagesRepository{
		db: db,
	}
}

func (r *PremiumPackagesRepository) ExistPackage(ctx context.Context, id uuid.UUID) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM premium_packages WHERE id = $1)`

	var exist bool
	err := r.db.QueryRowContext(ctx, query, id).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}
