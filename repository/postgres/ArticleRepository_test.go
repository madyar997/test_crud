package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"test_crud/models"
	"testing"
	"time"
)

func TestPostgresArticleRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	mockArticles := []models.Article{
		{
			Id:        1,
			AuthorId:  1,
			Title:     "test title 1",
			Body:      "test body 1",
			CreatedOn: time.Time{},
		},
		{
			Id:        2,
			AuthorId:  2,
			Title:     "test title 2",
			Body:      "test body 2",
			CreatedOn: time.Time{},
		},
	}
	rows := sqlmock.NewRows([]string{"id", "author_id", "title", "body", "created_on"}).
		AddRow(mockArticles[0].Id, mockArticles[0].AuthorId, mockArticles[0].Title, mockArticles[0].Body, mockArticles[0].CreatedOn).
		AddRow(mockArticles[1].Id, mockArticles[1].AuthorId, mockArticles[1].Title, mockArticles[1].Body, mockArticles[1].CreatedOn)
	query := `select id, title, author_id, body, created_on from articles`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewPostgresArticleRepository(db)
	articles, err := a.Get()
	assert.NoError(t, err)
	assert.NotNil(t, articles)
	assert.Len(t, articles, 2)
}
