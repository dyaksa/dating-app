package postgres

import (
	"context"
	"database/sql"

	"github.com/dyaksa/dating-app/internal/modules/swipes/v1/entity"
	"github.com/google/uuid"
)

type SwipesRepo struct {
	db *sql.DB
}

func NewSwipeRepository(db *sql.DB) *SwipesRepo {
	return &SwipesRepo{db: db}
}

func (r *SwipesRepo) Swipe(ctx context.Context, entity *entity.Swipes) error {
	query := `INSERT INTO swipes (swiper_id, target_profile_id, swipe_type, last_swipe_date, created_at) VALUES ($1, $2, $3, CURRENT_DATE, $4)`
	_, err := r.db.ExecContext(ctx, query, entity.SwiperID, entity.TargetProfileID, entity.SwipeType, entity.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *SwipesRepo) Count(ctx context.Context, userID uuid.UUID) (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM swipes WHERE swiper_id = $1 AND last_swipe_date = CURRENT_DATE`

	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
