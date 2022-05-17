package errs

type AppError struct {
	Code int    `json:"omitempty"`
	Msg  string `json:",msg"`
}

func NewAppError(msg string, code int) *AppError {
	return &AppError{
		code,
		msg,
	}
}
