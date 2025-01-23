package entity

type HttpError struct {
	StatusCode int
	Message string
}

func (err *HttpError) Error() string {
	return err.Message
}