package dto

import "twitter-uala/src/entities"

type UserIDHeader struct {
	UserID entities.UserID `header:"x-caller-id" binding:"required"`
}
