package validations

import "fmt"

type TweetTooLongError struct {
	Length int `json:"length"`
}

func (err TweetTooLongError) Error() string {
	return fmt.Sprintf("tweet's length (%d) exceeds maximum allowed", err.Length)
}

type EmptyTweetError struct{}

func (err EmptyTweetError) Error() string {
	return fmt.Sprint("tweet should not be empty")
}
