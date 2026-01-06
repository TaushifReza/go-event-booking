package common

type AppError struct {
	Code int
	Message error
}

func (e *AppError) Error() string{
	return e.Message.Error()
}