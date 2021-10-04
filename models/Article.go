package models

import (
	"time"
)

type Article struct {
	Id        int64     `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

type ArticleRepository interface {
	Get() (res []Article, err error)
	Create(a *Article) error
	Delete(id int64) error
}

type ArticleUsecase interface {
	Get() ([]Article, error)
	Create(a *Article) error
	Delete(id int64) error
}
