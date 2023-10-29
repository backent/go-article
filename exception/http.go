package exception

import (
	"net/http"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/web"
	"github.com/go-playground/validator/v10"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	var response web.WebResponse
	if err, ok := i.(ErrorNotFound); ok {
		response = web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found",
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
