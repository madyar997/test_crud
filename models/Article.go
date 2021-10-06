package models

import (
	"time"
)

type Article struct {
	Id        int64     `json:"id"`
	AuthorId  int64     `json:"author_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

type ArticleRepository interface {
	Get() (res []Article, err error)
	GetById(id int64) (a Article, err error)
	Create(a *Article) error
	Delete(id int64) error
	Update(ar *Article) error
	GetByAuthorId(authorId int64) (res []Article, err error)
}

type ArticleUsecase interface {
	Get() ([]Article, error)
	GetById(id int64) (Article, error)
	Create(a *Article) error
	Delete(id int64) error
	Update(ar *Article) error
	GetByAuthorId(authorId int64) ([]Article, error)
}
