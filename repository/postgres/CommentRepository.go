package postgres

import (
	"database/sql"
	"log"
	"test_crud/models"
)

type postgresCommentRepository struct {
	Conn *sql.DB
}

func NewPostgresCommentRepository(Conn *sql.DB) models.CommentRepository {
	return &postgresCommentRepository{Conn}
}

func (p *postgresCommentRepository) Get() (c []models.Comment, err error) {
	var comment models.Comment
	var comments []models.Comment
	query := `select id, content, article_id from comments`
	rows, err := p.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Content, &comment.ArticleId)
		if err != nil {
			log.Println("err while scanning form the row")
		}
		comments = append(comments, comment)
	}
	defer rows.Close()

	return comments, nil
}

func (p *postgresCommentRepository) Create(comment *models.Comment) error {
	query := `INSERT INTO "comments" (id, content, article_id) VALUES ($1, $2, $3)`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(comment.Id, comment.Content, comment.ArticleId)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresCommentRepository) Delete(id int64) error {
	query := `DELETE FROM "comments" WHERE id = $1`
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

func (p *postgresCommentRepository) Update(comment *models.Comment) error {
	query := `UPDATE "comments" set content=$1, article_id=$2 WHERE id = $3`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.Content, comment.ArticleId, comment.Id)
	if err != nil {
		return err
	}
	return nil
}
