package dto

import "twitter-uala/src/entities"

type TimelineResponse struct {
	TweetsCount int              `json:"tweets_count"`
	Tweets      []entities.Tweet `json:"tweets"`
}
