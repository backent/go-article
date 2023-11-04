package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/backent/go-article/helpers"
	"github.com/golang-jwt/jwt/v5"
)

type RepositoryAuthJWTImpl struct {
	secretKeys    []byte
	tokenLifeTime time.Time
}

func NewRepositoryAuthJWTImpl() RepositoryAuthInterface {

	tokenLifeTime, err := strconv.Atoi(os.Getenv("APP_TOKEN_EXPIRE_IN_SEC"))
	helpers.PanicIfError(err)

	return &RepositoryAuthJWTImpl{
		secretKeys:    []byte(os.Getenv("APP_SECRET_KEY")),
		tokenLifeTime: time.Now().Add(time.Second * time.Duration(tokenLifeTime)),
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
func (implementation *RepositoryAuthJWTImpl) Validate(tokenString string) bool {

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return implementation.secretKeys, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true
	} else {
		return false
	}
}
