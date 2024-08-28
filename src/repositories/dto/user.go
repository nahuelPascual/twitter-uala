package dto

import (
	"time"

	"twitter-uala/src/entities"
)

type User struct {
	ID        entities.UserID `db:"id"`
	Username  string          `db:"username"`
	CreatedAt time.Time       `db:"created_at"`
}

func (u User) ToDomain() entities.User {
	return entities.User{
		ID:        u.ID,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}
