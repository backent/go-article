package user

type UserRequestCreate struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UserRequestUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
