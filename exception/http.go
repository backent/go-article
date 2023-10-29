package exception

import (
	"net/http"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/web"
)

func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   i,
	}

	helpers.ReturnResponseJSON(w, response)
}
