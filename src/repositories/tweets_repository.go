package repositories

import (
	"github.com/jmoiron/sqlx"

	"twitter-uala/src/entities"
	"twitter-uala/src/repositories/dto"
)

type (
	TweetsRepository interface {
		Create(entities.Tweet) error
		GetTweets(_ entities.UserID, limit int) ([]entities.Tweet, error)
	}
	tweetsRepository struct {
		client *sqlx.DB
	}
)

func NewTweetsRepository(client *sqlx.DB) TweetsRepository {
	return &tweetsRepository{client: client}
}

func (r tweetsRepository) Create(tweet entities.Tweet) error {
	const query = "INSERT INTO tweets (content, user_id, created_at) VALUES (:content, :user_id, :created_at)"
	if _, err := r.client.NamedExec(query, dto.NewTweetDto(tweet)); err != nil {
		return err
	}
	return nil
}

func (r tweetsRepository) GetTweets(userID entities.UserID, limit int) ([]entities.Tweet, error) {
	const query = "SELECT * FROM tweets WHERE user_id = $1 ORDER BY created_at DESC LIMIT $2"
	var tweets dto.Tweets
	if err := r.client.Select(&tweets, query, userID, limit); err != nil {
		return nil, err
	}
	return tweets.ToDomain(), nil
}
