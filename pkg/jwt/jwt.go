package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Jwt struct {
	SecretKey []byte
}

func NewJwt(secretKey []byte) *Jwt {
	return &Jwt{SecretKey: secretKey}
}

func (j *Jwt) GenerateJWT(userIdentity string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Minute)
	claims["authorized"] = true
	claims["email"] = userIdentity

	tokenString, err := token.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
