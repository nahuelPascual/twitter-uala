package mocks

import "twitter-uala/src/entities"

type TweetRepositoryMock struct {
	CreateMock    func(tweet entities.Tweet) error
	GetTweetsMock func(_ entities.UserID, limit int) ([]entities.Tweet, error)
}

func (m TweetRepositoryMock) Create(tweet entities.Tweet) error {
	if m.CreateMock == nil {
		panic("CreateMock is nil")
	}
	return m.CreateMock(tweet)
}

func (m TweetRepositoryMock) GetTweets(userID entities.UserID, limit int) ([]entities.Tweet, error) {
	if m.GetTweetsMock == nil {
		panic("GetTweetsMock is nil")
	}
	return m.GetTweetsMock(userID, limit)
}
