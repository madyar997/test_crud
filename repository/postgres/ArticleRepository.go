package postgres

import (
	"database/sql"
	"log"
	"test_crud/models"
)

type postgresArticleRepository struct {
	Conn *sql.DB
}

func NewPostgresArticleRepository(Conn *sql.DB) models.ArticleRepository {
	return &postgresArticleRepository{Conn}
}

func (p *postgresArticleRepository) Get() (res []models.Article, err error) {
	var article = models.Article{}
	var articles []models.Article
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

func (p *postgresArticleRepository) Create(a *models.Article) error {
	query := `INSERT INTO "articles" (id, author, title, body, created_on) VALUES ($1, $2, $3, $4, $5)`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(a.Id, a.Author, a.Title, a.Body, a.CreatedOn)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	a.Id = lastID

	return nil
}

func (p *postgresArticleRepository) Delete(id int64) error {
	query := `DELETE FROM "articles" WHERE id = $1`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresArticleRepository) Update(ar *models.Article) error {
	query := `UPDATE "articles" set author=$1, title=$2, body=$3, created_on=$4 WHERE id = $5`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ar.Author, ar.Title, ar.Body, ar.CreatedOn, ar.Id)
	if err != nil {
		return err
	}
	return nil
}
