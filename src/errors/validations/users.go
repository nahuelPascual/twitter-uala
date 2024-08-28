package validations

type FollowingHimSelfError struct{}

func (err FollowingHimSelfError) Error() string {
	return "following him self is not allowed"
}
