package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	"twitter-uala/src/entities"
	"twitter-uala/src/repositories/dto"
)

type (
	TweetsRepository interface {
		Create(*gin.Context, entities.Tweet) error
	}
	tweetsRepository struct {
		client *sqlx.DB
	}
)

func NewTweetsRepository(client *sqlx.DB) TweetsRepository {
	return &tweetsRepository{client: client}
}

func (r tweetsRepository) Create(ctx *gin.Context, tweet entities.Tweet) error {
	const query = "INSERT INTO tweets (content, user_id, created_at) VALUES (:content, :user_id, :created_at)"
	if _, err := r.client.NamedExec(query, dto.NewTweetDto(tweet)); err != nil {
		return err
	}
	return nil
}
