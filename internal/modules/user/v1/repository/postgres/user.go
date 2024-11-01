package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/dyaksa/dating-app/internal/modules/user/v1/entity"
	"github.com/dyaksa/dating-app/internal/utils/password"
	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Save(ctx context.Context, entity *entity.User) error {
	query := `INSERT INTO users (name, email, password_hash, gender, birthdate, is_verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	passwordHash, err := password.Hash(entity.PasswordHash)
	if err != nil {
		return err
	}

	parseDate, err := time.Parse("2006-01-02", entity.BirthDate)
	if err != nil {
		return err
	}

	_, err = u.db.ExecContext(ctx, query, entity.Name, entity.Email, passwordHash, entity.Gender, parseDate, entity.IsVerified, entity.CreatedAt, entity.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) ExistUser(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, name, email, password_hash, created_at, updated_at FROM users WHERE email = $1`

	row := u.db.QueryRowContext(ctx, query, email)
	user := entity.User{}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) SwipeUser(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	query := `
	SELECT u.id, u.name, u.gender, u.birthdate, u.is_verified
	FROM users u
	LEFT JOIN swipes s ON u.id = s.target_profile_id
    AND s.swiper_id = $1
    AND s.last_swipe_date = CURRENT_DATE
	WHERE s.target_profile_id IS NULL
	AND u.id != $2
	ORDER BY RANDOM() 
	LIMIT 1`

	row := u.db.QueryRowContext(ctx, query, userID, userID)
	err := row.Scan(&user.ID, &user.Name, &user.Gender, &user.BirthDate, &user.IsVerified)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
