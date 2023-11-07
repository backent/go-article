package article

import (
	"github.com/backent/go-article/models"
	"github.com/backent/go-article/web/user"
)

type ArticleResponse struct {
	UserDetails user.UserResponse `json:"user_details"`
	UserId      int               `json:"user_id"`
	Title       string            `json:"title"`
	Content     string            `json:"content"`
}

func ArticleModelToArticleResponse(articleModel models.Article) ArticleResponse {
	return ArticleResponse{
		UserId:      articleModel.UserId,
		Title:       articleModel.Title,
		Content:     articleModel.Content,
		UserDetails: user.UserResponse{},
	}
}
