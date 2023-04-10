package helper

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

// hashedPassword is provided from repository record. providedPassword is provided by the users input.
// Return true if both password hash is match, or an error when no match.
func ComparePassword(hashedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	return err == nil
}
