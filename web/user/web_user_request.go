package user

type UserRequestCreate struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,max=72"`
}

type UserRequestUpdate struct {
	Id       int    `json:"id" validate:"required"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`

	// Not from request, for storing from middleware
	SavedPassword string
}

type UserRequestDelete struct {
	Id int
}
