package models

type User struct {
	Id       int
	Name     string
	Username string
	Password string
	Articles []Article
}
