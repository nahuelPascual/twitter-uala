package mocks

import (
	"twitter-uala/src/entities"
)

type UsersRepositoryMock struct {
	GetByUsernameMock    func(username string) (entities.User, error)
	AddFollowerMock      func(userID entities.UserID, followerID entities.UserID) error
	GetFollowedUsersMock func(id entities.UserID) ([]entities.UserID, error)
}

func (m UsersRepositoryMock) GetByUsername(username string) (entities.User, error) {
	if m.GetByUsernameMock == nil {
		panic("GetByUsernameMock is nil")
	}
	return m.GetByUsernameMock(username)
}

func (m UsersRepositoryMock) AddFollower(userID entities.UserID, followerID entities.UserID) error {
	if m.AddFollowerMock == nil {
		panic("AddFollowerMock is nil")
	}
	return m.AddFollowerMock(userID, followerID)
}

func (m UsersRepositoryMock) GetFollowedUsers(id entities.UserID) ([]entities.UserID, error) {
	if m.GetFollowedUsersMock == nil {
		panic("GetFollowedUsersMock is nil")
	}
	return m.GetFollowedUsersMock(id)
}
