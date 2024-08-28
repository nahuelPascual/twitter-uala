package services

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"twitter-uala/src/entities"
	"twitter-uala/src/repositories/mocks"

	"github.com/stretchr/testify/assert"
)

func Test_timelineService_ResolveTimeline(t *testing.T) {
	t.Run("fails when GetFollowedUsers returns an error", func(t *testing.T) {
		expectedErr := errors.New("an error")
		s := timelineService{
			UsersRepository: mocks.UsersRepositoryMock{
				GetFollowedUsersMock: func(entities.UserID) ([]entities.UserID, error) {
					return nil, expectedErr
				},
			},
		}

		tweets, gotErr := s.ResolveTimeline(1)

		assert.Equal(t, expectedErr, gotErr)
		assert.Empty(t, tweets)
	})
	t.Run("fails when GetTweets returns an error", func(t *testing.T) {
		expectedErr := errors.New("an error")
		s := timelineService{
			UsersRepository: mocks.UsersRepositoryMock{
				GetFollowedUsersMock: func(entities.UserID) ([]entities.UserID, error) {
					return []entities.UserID{1, 2, 3}, nil
				},
			},
			TweetsRepository: mocks.TweetRepositoryMock{
				GetTweetsMock: func(entities.UserID, int) ([]entities.Tweet, error) {
					return nil, expectedErr
				},
			},
		}

		tweets, gotErr := s.ResolveTimeline(1)

		assert.Equal(t, expectedErr, gotErr)
		assert.Empty(t, tweets)
	})
	t.Run("fails when GetTweets returns an error", func(t *testing.T) {
		s := timelineService{
			UsersRepository: mocks.UsersRepositoryMock{
				GetFollowedUsersMock: func(entities.UserID) ([]entities.UserID, error) {
					return []entities.UserID{1, 2, 3}, nil
				},
			},
			TweetsRepository: mocks.TweetRepositoryMock{
				GetTweetsMock: func(userID entities.UserID, limit int) ([]entities.Tweet, error) {
					tweets := make([]entities.Tweet, limit)
					for i := range tweets {
						tweets[i] = entities.Tweet{
							ID:        entities.TweetID(i + 1),
							Content:   fmt.Sprintf("tweet #%d", i),
							UserID:    userID,
							CreatedAt: time.Now(),
						}
					}
					return tweets, nil
				},
			},
		}

		tweets, gotErr := s.ResolveTimeline(1)

		assert.NoError(t, gotErr)
		assert.Equal(t, 500, len(tweets))
	})
}
