package services

import (
	"github.com/gin-gonic/gin"

	"twitter-uala/src/entities"
	"twitter-uala/src/repositories"
)

type (
	TweetsService interface {
		Publish(*gin.Context, entities.Tweet) error
	}
	tweetsService struct {
		TweetsRepository repositories.TweetsRepository
	}
)

func NewTweetsService(tweetsRepository repositories.TweetsRepository) TweetsService {
	return &tweetsService{TweetsRepository: tweetsRepository}
}

func (s tweetsService) Publish(ctx *gin.Context, tweet entities.Tweet) error {
	// TODO send creation event
	return s.TweetsRepository.Create(ctx, tweet)
}
