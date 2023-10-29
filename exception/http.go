package exception

import (
	"net/http"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/web"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	var response web.WebResponse
	errorNotFound, isErrorNotFound := i.(ErrorNotFound)
	if isErrorNotFound {
		response = web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Status Not Found",
			Data:   errorNotFound.Error(),
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
