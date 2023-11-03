package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RepositoryAuthJWTImpl struct {
	secretKeys    []byte
	tokenLifeTime time.Time
}

func NewRepositoryAuthJWTImpl() RepositoryAuthInterface {
	return &RepositoryAuthJWTImpl{
		secretKeys:    []byte("awfe8j32jin9"),
		tokenLifeTime: time.Now().Add(time.Hour * 3),
	}
}

func (implementation *RepositoryAuthJWTImpl) Issue(payload string) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(implementation.tokenLifeTime),
		Issuer:    payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generatedToken, err := token.SignedString(implementation.secretKeys)
	return generatedToken, err
}
func (implementation *RepositoryAuthJWTImpl) Validate(token string) (string, error) {

	panic("awd")
}
