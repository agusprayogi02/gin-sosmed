package helper

import (
	"time"

	"gin-sosmed/entity"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *entity.User, signingKeys string) (string, error) {
	claims := JWTClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signing, err := HashPassword(signingKeys)
	if err != nil {
		return signing, err
	}
	return token.SignedString([]byte(signing))
}
