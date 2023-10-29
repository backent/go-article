package user

import "github.com/backent/go-article/models"

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func UserModelToResponse(user models.User) UserResponse {
	return UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Name:     user.Name,
	}
}

func UsersModelToResponses(users []models.User) []UserResponse {
	var responses []UserResponse
	for _, val := range users {
		responses = append(responses, UserModelToResponse(val))
	}

	return responses
}
