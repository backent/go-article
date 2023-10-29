package user

type UserRequestCreate struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

type UserRequestUpdate struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
