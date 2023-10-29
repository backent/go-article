package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/backent/go-article/helpers"
	services "github.com/backent/go-article/services/user"
	"github.com/backent/go-article/web"
	webUser "github.com/backent/go-article/web/user"
	"github.com/julienschmidt/httprouter"
)

type ControllerUserImpl struct {
	userService services.ServiceUserInterface
}

func NewControllerUser(userService services.ServiceUserInterface) ControllerUserInterface {
	return &ControllerUserImpl{
		userService: userService,
	}
}
func (implementation *ControllerUserImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userCreateRequest := webUser.UserRequestCreate{}
	helpers.DecodeRequestBody(r, &userCreateRequest)

	ctx := context.Background()

	user := implementation.userService.Create(ctx, userCreateRequest)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   user,
	}

	helpers.ReturnResponseJSON(w, response)
}
func (implementation *ControllerUserImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userUpdateRequest := webUser.UserRequestUpdate{}
	helpers.DecodeRequestBody(r, &userUpdateRequest)

	ctx := context.Background()
	id, err := strconv.Atoi(param.ByName("id"))
	helpers.PanicIfError(err)

	userUpdateRequest.Id = id
	user := implementation.userService.Update(ctx, userUpdateRequest)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   user,
	}

	helpers.ReturnResponseJSON(w, response)

}
func (implementation *ControllerUserImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	ctx := context.Background()
	id, err := strconv.Atoi(param.ByName("id"))
	helpers.PanicIfError(err)

	implementation.userService.Delete(ctx, id)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   nil,
	}

	helpers.ReturnResponseJSON(w, response)
}
func (implementation *ControllerUserImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	id, err := strconv.Atoi(param.ByName("id"))
	helpers.PanicIfError(err)

	user := implementation.userService.FindById(ctx, id)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   user,
	}

	helpers.ReturnResponseJSON(w, response)

}
func (implementation *ControllerUserImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	ctx := context.Background()
	users := implementation.userService.FindAll(ctx)

	response := web.WebResponse{
		Status: "OK",
		Code:   http.StatusOK,
		Data:   users,
	}

	helpers.ReturnResponseJSON(w, response)
}
