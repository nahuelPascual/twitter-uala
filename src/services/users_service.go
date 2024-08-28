package services

import (
	"twitter-uala/src/entities"
	"twitter-uala/src/errors/validations"
	"twitter-uala/src/repositories"
)

type (
	UsersService interface {
		Follow(callerID entities.UserID, followedUser string) error
	}
	usersService struct {
		UsersRepository repositories.UsersRepository
	}
)

func NewUsersService(usersRepository repositories.UsersRepository) UsersService {
	return usersService{UsersRepository: usersRepository}
}

func (s usersService) Follow(callerID entities.UserID, followedUser string) error {
	// check existence
	user, err := s.UsersRepository.GetByUsername(followedUser)
	if err != nil {
		return err
	}

	if callerID == user.ID {
		return validations.FollowingHimSelfError{}
	}

	return s.UsersRepository.AddFollower(user.ID, callerID)
}
