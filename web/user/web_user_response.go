package user

import "github.com/backent/go-article/models"

type UserArticleResponse struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type UserResponse struct {
	Id       int                   `json:"id"`
	Username string                `json:"username"`
	Name     string                `json:"name"`
	Articles []UserArticleResponse `json:"articles"`
}

func UserModelToResponse(user models.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
		Articles: UserArticleModelToUserArticleResponse(user.Articles),
	}
}

func UsersModelToResponses(users []models.User) []UserResponse {
	var responses []UserResponse
	for _, val := range users {
		responses = append(responses, UserModelToResponse(val))
	}

	return responses
}

func UserArticleModelToUserArticleResponse(articles []models.Article) []UserArticleResponse {
	var userArticleReponse []UserArticleResponse
	if len(articles) > 0 {
		for _, val := range articles {
			userArticleReponse = append(userArticleReponse, UserArticleResponse{
				Id:    val.Id,
				Title: val.Title,
			})
		}
		return userArticleReponse
	} else {
		return make([]UserArticleResponse, 0)
	}
}
