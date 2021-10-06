package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"test_crud/models"
	"testing"
)

func TestPostgresAuthorRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	mockAuthors := []models.Author{
		{
			Id:        1,
			FirstName: "Madyar",
			LastName:  "Turgenbayev",
			Email:     "madiar.997@gmail.com",
		},
		{
			Id:        2,
			FirstName: "Maya",
			LastName:  "Turgenbayeva",
			Email:     "maya.turgenbaev@mail.ru",
		},
	}
	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "email"}).
		AddRow(mockAuthors[0].Id, mockAuthors[0].FirstName, mockAuthors[0].LastName, mockAuthors[0].Email).
		AddRow(mockAuthors[1].Id, mockAuthors[1].FirstName, mockAuthors[1].LastName, mockAuthors[1].Email)
	query := `SELECT id, first_name, last_name, email FROM authors`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewPostgresAuthorRepository(db)
	authors, err := a.Get()
	assert.NoError(t, err)
	assert.NotNil(t, authors)
	assert.Len(t, authors, 2)
}

func TestPostgresAuthorRepository_Create(t *testing.T) {
	author := &models.Author{
		Id:        1,
		FirstName: "Madyar",
		LastName:  "Turgenbayev",
		Email:     "madiar.997@gmail.com",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "INSERT INTO \"authors\" \\(id, first_name, last_name, email\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(author.Id, author.FirstName, author.LastName, author.Email).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresAuthorRepository(db)
	err = a.Create(author)
	assert.NoError(t, err)
}

func TestPostgresArticleRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	query := "DELETE FROM \"authors\" WHERE id = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewPostgresAuthorRepository(db)
	id := int64(1)
	err = a.Delete(id)
	assert.NoError(t, err)
}

func TestPostgresAuthorRepository_Update(t *testing.T) {
	author := &models.Author{
		Id:        1,
		FirstName: "Madyar",
		LastName:  "Turgenbayev",
		Email:     "madiar.997@gmail.com",
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "UPDATE \"authors\" set first_name=\\$1, last_name=\\$2, email=\\$3 WHERE id = \\$4"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(author.FirstName, author.LastName, author.Email, author.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresAuthorRepository(db)
	err = a.Update(author)
	assert.NoError(t, err)
}
