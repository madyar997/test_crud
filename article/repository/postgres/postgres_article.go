package postgres

import (
	"database/sql"
	"log"
	"test_crud/models"
)

type postgresArticleRepository struct {
	Conn *sql.DB
}

func NewPostgresArticleRepository(Conn *sql.DB) models.ArticleRepository{
	return &postgresArticleRepository{Conn}
}

func (p *postgresArticleRepository) Get() (res []models.Article, err error) {
	var article = models.Article{}
	var articles []models.Article
	//define query string
	query := `select id, title, author, body, created_on from articles`
	rows, err := p.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.Author, &article.Body, &article.CreatedOn)
		if err != nil {
			log.Println("err while scanning form the row")
		}
		articles = append(articles, article)
	}
	defer rows.Close()

	return articles, nil
}