package postgres

import (
	"database/sql"
	"log"
	"test_crud/models"
)

type postgresAuthorRepository struct {
	Conn *sql.DB
}

func NewPostgresAuthorRepository(Conn *sql.DB) models.AuthorRepository {
	return &postgresAuthorRepository{Conn}
}

func (p *postgresAuthorRepository) Get() (a []models.Author, err error) {
	var author models.Author
	var authors []models.Author
	query := `SELECT id, first_name, last_name, email FROM authors`
	rows, err := p.Conn.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&author.Id, &author.FirstName, &author.LastName, &author.Email)
		if err != nil {
			log.Println("err while scanning form the row")
		}
		authors = append(authors, author)
	}
	defer rows.Close()

	return authors, nil
}

func (p *postgresAuthorRepository) Create(a *models.Author) error {
	query := `INSERT INTO "authors" (id, first_name, last_name, email) VALUES ($1, $2, $3, $4)`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(a.Id, a.FirstName, a.LastName, a.Email)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresAuthorRepository) Delete(id int64) error {
	query := `DELETE FROM "authors" WHERE id = $1`
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

func (p *postgresAuthorRepository) Update(a *models.Author) error {
	query := `UPDATE "authors" set first_name=$1, last_name=$2, email=$3 WHERE id = $4`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(a.FirstName, a.LastName, a.Email, a.Id)
	if err != nil {
		return err
	}
	return nil
}
