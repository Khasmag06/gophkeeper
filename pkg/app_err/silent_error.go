package app_err

type SilentError struct {
	message string
}

func NewSilentError(msg string) error {
	return SilentError{
		message: msg,
	}
}
