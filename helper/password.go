package helper

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	//PanicIfError(err)
	return string(hashedPassword), err
}

func ComparePassword(hashedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(providedPassword))
	Error.Println(err)
	return err == nil
}
