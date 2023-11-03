package auth

import (
	"context"
	"net/http"

	"github.com/backent/go-article/helpers"
	servicesAuth "github.com/backent/go-article/services/auth"
	"github.com/backent/go-article/web"
	webAuth "github.com/backent/go-article/web/auth"
	"github.com/julienschmidt/httprouter"
)

type ControllerAuthImpl struct {
	servicesAuth.ServiceAuthInterface
}

func NewControllerAuthImpl(servicesAuth servicesAuth.ServiceAuthInterface) ControllerAuthInterface {
	return &ControllerAuthImpl{ServiceAuthInterface: servicesAuth}
}

func (implementation *ControllerAuthImpl) Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	var loginRequest webAuth.LoginRequest

	helpers.DecodeRequestBody(r, &loginRequest)

	ctx := context.Background()

	loginResponse := implementation.ServiceAuthInterface.Login(ctx, loginRequest)

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   loginResponse,
	}

	helpers.ReturnResponseJSON(w, response)
}
