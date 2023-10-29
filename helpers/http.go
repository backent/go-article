package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/backent/go-article/web"
)

func DecodeRequestBody(r *http.Request, any interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(any)
	PanicIfError(err)
}

func ReturnResponseJSON(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
