package repositories

import (
	"database/sql"
	"errors"

	"twitter-uala/src/entities"
	apierrors "twitter-uala/src/errors"
	"twitter-uala/src/repositories/dto"

	"github.com/jmoiron/sqlx"
)

type (
	UsersRepository interface {
		GetByUsername(username string) (entities.User, error)
		AddFollower(userID entities.UserID, followerID entities.UserID) error
		GetFollowedUsers(entities.UserID) ([]entities.UserID, error)
	}
	usersRepository struct {
		client *sqlx.DB
	}
)

func NewUsersRepository(db *sqlx.DB) UsersRepository {
	return &usersRepository{db}
}

func (r usersRepository) GetByUsername(username string) (entities.User, error) {
	const query = `SELECT * FROM users WHERE username = $1`
	var user dto.User
	if err := r.client.Get(&user, query, username); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.User{}, apierrors.EntityNotFoundError{Entity: "user", Key: "username", Value: username}
		}
		return entities.User{}, err
	}

	return user.ToDomain(), nil
}

func (r usersRepository) AddFollower(userID entities.UserID, followerID entities.UserID) error {
	const query = "INSERT INTO followers (user_id, follower_user_id) VALUES ($1, $2)"
	if _, err := r.client.Exec(query, userID, followerID); err != nil {
		return err
	}
	return nil
}

func (r usersRepository) GetFollowedUsers(userID entities.UserID) ([]entities.UserID, error) {
	const query = `SELECT user_id FROM followers WHERE follower_user_id = $1`
	var userIDs []entities.UserID
	if err := r.client.Select(&userIDs, query, userID); err != nil {
		return nil, err
	}
	return userIDs, nil
}
