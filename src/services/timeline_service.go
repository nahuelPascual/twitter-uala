package services

import (
	"math"
	"sort"

	"twitter-uala/src/entities"
	"twitter-uala/src/repositories"
)

type (
	TimelineService interface {
		ResolveTimeline(entities.UserID) ([]entities.Tweet, error)
	}
	timelineService struct {
		UsersRepository  repositories.UsersRepository
		TweetsRepository repositories.TweetsRepository
	}
)

const timelineLimit = 500

func NewTimelineService(usersRepository repositories.UsersRepository, tweetsRepository repositories.TweetsRepository) timelineService {
	return timelineService{
		UsersRepository:  usersRepository,
		TweetsRepository: tweetsRepository,
	}
}

func (s timelineService) ResolveTimeline(callerID entities.UserID) ([]entities.Tweet, error) {
	followedUsers, err := s.UsersRepository.GetFollowedUsers(callerID)
	if err != nil {
		return nil, err
	}

	resultsChan := make(chan []entities.Tweet, len(followedUsers))
	errChan := make(chan error, len(followedUsers))
	for _, userID := range followedUsers {
		go func(userID entities.UserID) {
			tweets, err := s.TweetsRepository.GetTweets(userID, timelineLimit)
			if err != nil {
				errChan <- err
			} else {
				resultsChan <- tweets
			}
		}(userID)
	}

	var timeline []entities.Tweet
	for range followedUsers {
		select {
		case err = <-errChan:
			return nil, err
		case tweets := <-resultsChan:
			timeline = append(timeline, tweets...)
		}
	}

	// TODO it could be expensive and should be improved
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].CreatedAt.After(timeline[j].CreatedAt)
	})

	maxSize := int(math.Min(float64(len(timeline)), float64(timelineLimit)))
	return timeline[:maxSize], nil
}
