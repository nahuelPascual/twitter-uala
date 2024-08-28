package entities

import "time"

type (
	UserID uint64
	User   struct {
		ID        UserID    `json:"id"`
		Username  string    `json:"username"`
		Followers []UserID  `json:"followers"`
		Following []UserID  `json:"following"`
		CreatedAt time.Time `json:"created_at"`
	}
)
