package services

import (
	"twitter-uala/src/entities"
	"twitter-uala/src/repositories"
)

type (
	TweetsService interface {
		Publish(entities.Tweet) error
	}
	tweetsService struct {
		TweetsRepository repositories.TweetsRepository
	}
)

func NewTweetsService(tweetsRepository repositories.TweetsRepository) TweetsService {
	return &tweetsService{TweetsRepository: tweetsRepository}
}

func (s tweetsService) Publish(tweet entities.Tweet) error {
	// TODO send creation event
	return s.TweetsRepository.Create(tweet)
}
