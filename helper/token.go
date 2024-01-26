package helper

import (
	"errors"
	"strings"
	"time"

	"gin-sosmed/entity"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTClaims struct {
	ID uuid.UUID `json:"id"`
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
	arrKey := strings.Split(signingKeys, "-")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(arrKey[0]))
	if err != nil {
		return "", err
	}

	return Encrypt(t, arrKey[len(arrKey)-1])
}

func VerifyToken(token string, signingKeys string) (*uuid.UUID, error) {
	arrKey := strings.Split(signingKeys, "-")
	tokenStr, err := Decrypt(token, arrKey[len(arrKey)-1])
	if err != nil {
		return nil, errors.New("signature has invalid")
	}
	realToken, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(arrKey[0]), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("signature has invalid")
		}
		return nil, errors.New("your token was expired")
	}

	claims, ok := realToken.Claims.(*JWTClaims)
	if !ok || !realToken.Valid {
		return nil, errors.New("your token was expired or invalid")
	}

	return &claims.ID, nil
}
