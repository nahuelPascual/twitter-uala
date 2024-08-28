package dto

import (
	"time"

	"twitter-uala/src/entities"
)

type (
	Tweet struct {
		ID        entities.TweetID `db:"id"`
		Content   string           `db:"content"`
		UserID    entities.UserID  `db:"user_id"`
		CreatedAt time.Time        `db:"created_at"`
	}
	Tweets []Tweet
)

func NewTweetDto(tweet entities.Tweet) Tweet {
	return Tweet{
		Content:   tweet.Content,
		UserID:    tweet.UserID,
		CreatedAt: tweet.CreatedAt,
	}
}

func (t Tweets) ToDomain() []entities.Tweet {
	tweets := make([]entities.Tweet, len(t))
	for i, tweet := range t {
		tweets[i] = entities.Tweet{
			ID:        tweet.ID,
			Content:   tweet.Content,
			UserID:    tweet.UserID,
			CreatedAt: tweet.CreatedAt,
		}
	}
	return tweets
}
