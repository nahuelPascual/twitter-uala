package dto

import (
	"time"

	"twitter-uala/src/entities"
)

type Tweet struct {
	ID        entities.TweetID `db:"id"`
	Content   string           `db:"content"`
	UserID    entities.UserID  `db:"user_id"`
	CreatedAt time.Time        `db:"created_at"`
}

func NewTweetDto(tweet entities.Tweet) Tweet {
	return Tweet{
		Content:   tweet.Content,
		UserID:    tweet.UserID,
		CreatedAt: tweet.CreatedAt,
	}
}
