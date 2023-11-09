package article

import (
	"github.com/backent/go-article/models"
)

type ArticleUserReponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type ArticleResponse struct {
	UserDetails ArticleUserReponse `json:"user_details"`
	UserId      int                `json:"user_id"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
}

func ArticleModelToResponse(articleModel models.Article) ArticleResponse {
	return ArticleResponse{
		UserId:  articleModel.UserId,
		Title:   articleModel.Title,
		Content: articleModel.Content,
		UserDetails: ArticleUserReponse{
			Id:       articleModel.UserDetails.Id,
			Username: articleModel.UserDetails.Username,
		},
	}
}

func ArticlesModelToResponses(articleModel []models.Article) []ArticleResponse {
	var articlesResponses []ArticleResponse
	for _, article := range articleModel {
		articlesResponses = append(articlesResponses, ArticleModelToResponse(article))
	}
	return articlesResponses
}
