package entity

import (
	"time"

	"github.com/google/uuid"
)

type Swipes struct {
	ID              uuid.UUID `db:"id"`
	SwiperID        uuid.UUID `db:"swiper_id"`
	TargetProfileID uuid.UUID `db:"target_profile_id"`
	SwipeType       bool      `db:"swipe_type"`
	CreatedAt       time.Time `db:"created_at"`
}
