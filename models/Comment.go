package models

type Comment struct {
	Id        int64  `json:"id"`
	Content   string `json:"content"`
	ArticleId int64  `json:"article_id"`
}

type CommentRepository interface {
	Get() (a []Comment, err error)
	Create(a *Comment) error
	Delete(id int64) error
	Update(a *Comment) error
}

type CommentUsecase interface {
	Get() ([]Comment, error)
	Create(a *Comment) error
	Update(a *Comment) error
	Delete(id int64) error
}
