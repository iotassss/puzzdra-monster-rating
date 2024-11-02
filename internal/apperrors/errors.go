package apperrors

type ErrValidation struct {
	Message string
	Cause   error
}

func (e *ErrValidation) Error() string {
	return e.Message
}

type ErrNotFound struct {
	Message string
	Cause   error
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

