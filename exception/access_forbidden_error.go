package exception

type AccessForbiddenError struct {
	Error string
}

func NewAccessForbidden(error string) AuthError {
	return AuthError{Error: error}
}
