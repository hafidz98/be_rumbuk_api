package helper

import (
	"fmt"
	"os"

	//"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hafidz98/be_rumbuk_api/models/service"
)

type CustomClaims struct {
	//UserData interface{} `json:"user_data"`
	UserData *service.GlobalJWTResponse `json:"user_data"`
	jwt.RegisteredClaims
}

func GenerateJWT(userData *service.GlobalJWTResponse, claims jwt.RegisteredClaims) (tokenString string, err error) {
	var jwtKey = []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
	fmt.Printf("key: \t %x \n", jwtKey)
	//expTime := time.Now().Add(30 * time.Minute)

	claim := &CustomClaims{
		UserData:         userData,
		RegisteredClaims: claims,
	}

	// claims := &CustomClaims{
	// 	UserData: userData,
	// 	RegisteredClaims: jwt.RegisteredClaims{
	// 		ExpiresAt: jwt.NewNumericDate(expTime),
	// 		IssuedAt:  jwt.NewNumericDate(time.Now()),
	// 	},
	// }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	var jwtKey = []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if _, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return
	}
	return
}

func ExtractRoleFromToken(signedToken string) (*service.GlobalJWTResponse, error) {
	var jwtKey = []byte(os.Getenv("JWT_ACCESS_SECRET_KEY"))
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claim.UserData, nil
	}
	return nil, err
}
