package helpers

type UserNotFoundError struct {
	Message string `default:"User not found."`
}

func (e *UserNotFoundError) Error() string {
	return e.Message
}
