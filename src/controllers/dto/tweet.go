package dto

import (
	"time"

	"twitter-uala/src/entities"
	"twitter-uala/src/errors/validations"
)

type Tweet struct {
	Content string `json:"content"`
}

func (t Tweet) Validate() error {
	const maxLenAllowed = 280
	tweetLen := len(t.Content)
	if tweetLen == 0 {
		return validations.EmptyTweetError{}
	}

	if tweetLen > maxLenAllowed {
		return validations.TweetTooLongError{Length: tweetLen}
	}

	return nil
}

func (t Tweet) ToDomain(userID entities.UserID) entities.Tweet {
	return entities.Tweet{
		Content:   t.Content,
		UserID:    userID,
		CreatedAt: time.Now().UTC(),
	}
}
