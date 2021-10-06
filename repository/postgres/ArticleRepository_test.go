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
	currentTime := time.Now()
	mockArticles := []models.Article{
		{
			Id:        1,
			AuthorId:  1,
			Title:     "test title 1",
			Body:      "test body 1",
			CreatedOn: currentTime,
		},
		{
			Id:        2,
			AuthorId:  2,
			Title:     "test title 2",
			Body:      "test body 2",
			CreatedOn: currentTime,
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

func TestPostgresArticleRepository_Create(t *testing.T) {
	currentTime := time.Now()
	article := &models.Article{
		Id:        1,
		AuthorId:  1,
		Title:     "test title 1",
		Body:      "test body 1",
		CreatedOn: currentTime,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "INSERT INTO \"articles\" VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(article.Id, article.AuthorId, article.Title, article.Body, article.CreatedOn).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresArticleRepository(db)
	err = a.Create(article)
	assert.NoError(t, err)
}

func TestPostgresAuthorRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	query := "DELETE FROM \"articles\" WHERE id \\= \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewPostgresArticleRepository(db)
	id := int64(1)
	err = a.Delete(id)
	assert.NoError(t, err)
}

func TestPostgresArticleRepository_Update(t *testing.T) {
	currentTime := time.Now()
	article := &models.Article{
		Id:        1,
		AuthorId:  1,
		Title:     "test title 1",
		Body:      "test body 1",
		CreatedOn: currentTime,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "UPDATE \"articles\" set author_id=\\$1, title=\\$2, body=\\$3, created_on=\\$4 WHERE id = \\$5"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(article.AuthorId, article.Title, article.Body, article.CreatedOn, article.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresArticleRepository(db)
	err = a.Update(article)
	assert.NoError(t, err)
}

func TestPostgresArticleRepository_GetById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	currentTime := time.Now()
	expectedArticle := models.Article{
		Id:        1,
		AuthorId:  1,
		Title:     "test title 1",
		Body:      "test body 1",
		CreatedOn: currentTime,
	}
	row := sqlmock.NewRows([]string{"id", "author_id", "title", "body", "created_on"}).
		AddRow(expectedArticle.Id, expectedArticle.AuthorId, expectedArticle.Title, expectedArticle.Body, expectedArticle.CreatedOn)
	query := "SELECT * from \"articles\" WHERE id = \\$1"
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WillReturnRows(row)
	a := NewPostgresArticleRepository(db)
	id := int64(1)
	actualArticle, err := a.GetById(id)
	//assert.Equal(t, expectedArticle, actualArticle)
	assert.NotNil(t, actualArticle)
	assert.NoError(t, err)
}
