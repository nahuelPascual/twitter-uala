package entities

import "time"

type (
	TweetID uint64
	Tweet   struct {
		ID        TweetID   `json:"id"`
		Content   string    `json:"content"`
		UserID    UserID    `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}
)
