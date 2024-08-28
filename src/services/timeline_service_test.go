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
	t.Run("success and ordering fine", func(t *testing.T) {
		now := time.Now().UTC()
		oldestTweet := entities.Tweet{
			ID:        1,
			Content:   "Tuit 1",
			UserID:    1,
			CreatedAt: now,
		}
		secondTweet := entities.Tweet{
			ID:        2,
			Content:   "Tuit 1",
			UserID:    1,
			CreatedAt: now.Add(1 * time.Hour),
		}
		newestTweet := entities.Tweet{
			ID:        3,
			Content:   "Tuit 1",
			UserID:    1,
			CreatedAt: now.Add(2 * time.Hour),
		}
		s := timelineService{
			UsersRepository: mocks.UsersRepositoryMock{
				GetFollowedUsersMock: func(entities.UserID) ([]entities.UserID, error) {
					return []entities.UserID{1}, nil
				},
			},
			TweetsRepository: mocks.TweetRepositoryMock{
				GetTweetsMock: func(userID entities.UserID, limit int) ([]entities.Tweet, error) {
					return []entities.Tweet{secondTweet, oldestTweet, newestTweet}, nil
				},
			},
		}

		tweets, gotErr := s.ResolveTimeline(2)

		assert.NoError(t, gotErr)
		assert.Equal(t, newestTweet, tweets[0])
		assert.Equal(t, secondTweet, tweets[1])
		assert.Equal(t, oldestTweet, tweets[2])
	})
	t.Run("success returning max allowed", func(t *testing.T) {
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
