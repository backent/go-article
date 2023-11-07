package article

import (
	"context"
	"net/http"

	"github.com/backent/go-article/helpers"
	servicesArticle "github.com/backent/go-article/services/article"
	"github.com/backent/go-article/web"
	webArticle "github.com/backent/go-article/web/article"
	"github.com/julienschmidt/httprouter"
)

type ControllerArticleImpl struct {
	servicesArticle.ServicesArticleInterface
}

func NewControllerArticleImpl(servicesArticle servicesArticle.ServicesArticleInterface) ControllerArticleInterface {
	return &ControllerArticleImpl{
		ServicesArticleInterface: servicesArticle,
	}
}

func (implementation *ControllerArticleImpl) Create(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	articleCreateRequest := webArticle.ArticleRequestCreate{}
	helpers.DecodeRequestBody(r, &articleCreateRequest)

	ctx := context.WithValue(r.Context(), helpers.ContextKey("token"), r.Header.Get("Authorization"))
	response := implementation.ServicesArticleInterface.Create(ctx, articleCreateRequest)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}

	helpers.ReturnResponseJSON(w, webResponse)

}
func (implementation *ControllerArticleImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	panic("implement me")
}
func (implementation *ControllerArticleImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	panic("implement me")
}
func (implementation *ControllerArticleImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	panic("implement me")
}
func (implementation *ControllerArticleImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	panic("implement me")
}
