package jwt

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type Interface interface {
	CreateJWTToken(UserId string) (string, error)
	VerifyJWTToken(tokenString string) (string, error)
}

type jsonWebToken struct {
	SecretKey string
}

func Init() Interface {
	secretKey := os.Getenv("JWT_SECRET")

	return &jsonWebToken{SecretKey: secretKey}
} 

func (j *jsonWebToken) CreateJWTToken(UserId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId": UserId,
	})

	tokenString, err := token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jsonWebToken) VerifyJWTToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error){
		return []byte(j.SecretKey), nil
	})
	
	if err != nil {
		return "", err
	}


	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims["UserId"].(string), nil
	}
	return "", err

}
