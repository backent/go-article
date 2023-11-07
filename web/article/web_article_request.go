package article

type ArticleRequestCreate struct {
	UserId  int    `json:"user_id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type ArticleRequestUpdate struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

type ArticleRequestDelete struct {
	Id int
}

type ArticleRequestFindById struct {
	Id int
}

type ArticleRequestFindAll struct{}
