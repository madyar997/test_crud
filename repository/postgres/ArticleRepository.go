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
	query := `select id, title, author_id, body, created_on from articles`
	rows, err := p.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&article.Id, &article.Title, &article.AuthorId, &article.Body, &article.CreatedOn)
		if err != nil {
			log.Println("err while scanning form the row")
		}
		articles = append(articles, article)
	}
	defer rows.Close()

	return articles, nil
}

func (p *postgresArticleRepository) Create(a *models.Article) error {
	query := `INSERT INTO "articles" VALUES ($1, $2, $3, $4, $5)`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Id, a.AuthorId, a.Title, a.Body, a.CreatedOn)
	if err != nil {
		return err
	}
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
	query := `UPDATE "articles" set author_id=$1, title=$2, body=$3, created_on=$4 WHERE id = $5`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ar.AuthorId, ar.Title, ar.Body, ar.CreatedOn, ar.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresArticleRepository) GetByAuthorId(authorId int64) (res []models.Article, err error) {
	var article = models.Article{}
	var articleList []models.Article
	query := `SELECT * from "articles" WHERE author_id = $1`
	rows, err := p.Conn.Query(query, authorId)
	for rows.Next() {
		rows.Scan(&article.Id, &article.AuthorId, &article.Title, &article.Body, &article.CreatedOn)
		if err != nil {
			log.Fatal("error while scanning the raws", err)
		}
		articleList = append(articleList, article)
	}
	return articleList, nil
}

func (p *postgresArticleRepository) GetById(id int64) (a models.Article, err error) {
	var article = models.Article{}
	query := `SELECT * from "articles" WHERE id = $1`
	row := p.Conn.QueryRow(query, id)
	err = row.Scan(&article.Id, &article.AuthorId, &article.Title, &article.Body, &article.CreatedOn)
	if err != nil {
		log.Println("err while scanning form the row")
	}

	return article, nil
}
