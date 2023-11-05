package middlewares

import (
	"context"

	"github.com/backent/go-article/exception"
	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/repositories/auth"
)

func ValidateToken(ctx context.Context, repositoriesAuth auth.RepositoryAuthInterface) int {
	var tokenString string
	token := ctx.Value(helpers.ContextKey("token"))
	tokenString, ok := token.(string)
	if !ok || tokenString == "" {
		helpers.PanicIfError(exception.NewUnAuthorized(""))
	}

	idInt, isValid := repositoriesAuth.Validate(tokenString)
	if !isValid {
		helpers.PanicIfError(exception.NewUnAuthorized(""))
	}
	return idInt
}
