package dto

type FollowUserParams struct {
	Username string `uri:"username,string" binding:"required"`
}
