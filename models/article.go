package models

import (
	"time"
)

type Article struct {
	Id uint32 `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}

type ArticleRepository interface {
	Get()  []Article
}

type ArticleUsecase interface {
	GET() []Article
}
