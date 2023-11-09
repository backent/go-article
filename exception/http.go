package exception

import (
	"net/http"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/web"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	logger := logrus.New()
	logger.Warn(i)
	var response web.WebResponse
	if err, ok := i.(ErrorNotFound); ok {
		response = web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   err.Error(),
		}
	} else if err, ok := i.(Unauthorized); ok {
		response = web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   err.Error(),
		}
	} else if err, ok := i.(BadRequest); ok {
		response = web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
	} else if err, ok := i.(validator.ValidationErrors); ok {
		response = web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
	} else if err, ok := i.(error); ok {
		response = web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Interal Server Error",
			Data:   err.Error(),
		}
	} else {
		response = web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   i,
		}
	}

	helpers.ReturnResponseJSON(w, response)
}
