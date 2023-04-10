package exception

type AuthError struct {
	Error string
}

func NewAuthorization(error string) AuthError {
	return AuthError{Error: error}
}
