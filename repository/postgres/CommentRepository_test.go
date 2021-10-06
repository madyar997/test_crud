package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"test_crud/models"
	"testing"
)

func TestPostgresCommentRepository_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	mockComments := []models.Comment{
		{
			Id:        1,
			Content:   "Comment content test1",
			ArticleId: 1,
		},
		{
			Id:        2,
			Content:   "Comment content test2",
			ArticleId: 2,
		},
	}
	rows := sqlmock.NewRows([]string{"id", "content", "article_id"}).
		AddRow(mockComments[0].Id, mockComments[0].Content, mockComments[0].ArticleId).
		AddRow(mockComments[1].Id, mockComments[1].Content, mockComments[1].ArticleId)
	query := `select id, content, article_id from comments`
	mock.ExpectQuery(query).WillReturnRows(rows)
	a := NewPostgresCommentRepository(db)
	comments, err := a.Get()
	assert.NoError(t, err)
	assert.NotNil(t, comments)
	assert.Len(t, comments, 2)
}

func TestPostgresCommentRepository_Create(t *testing.T) {
	comment := &models.Comment{
		Id:        1,
		Content:   "Comment content test1",
		ArticleId: 1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "INSERT INTO \"comments\" \\(id, content, article_id\\) VALUES \\(\\$1, \\$2, \\$3\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(comment.Id, comment.Content, comment.ArticleId).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresCommentRepository(db)
	err = a.Create(comment)
	assert.NoError(t, err)
}

func TestPostgresCommentRepository_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}

	query := "DELETE FROM \"comments\" WHERE id = \\$1"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewPostgresCommentRepository(db)
	id := int64(1)
	err = a.Delete(id)
	assert.NoError(t, err)
}

func TestPostgresCommentRepository_Update(t *testing.T) {
	comment := &models.Comment{
		Id:        1,
		Content:   "Comment content test1",
		ArticleId: 1,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("an error occured while opening a stub db connection :", err)
	}
	query := "UPDATE \"comments\" set content=\\$1, article_id=\\$2 WHERE id = \\$3"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(comment.Content, comment.ArticleId, comment.Id).WillReturnResult(sqlmock.NewResult(1, 1))
	a := NewPostgresCommentRepository(db)
	err = a.Update(comment)
	assert.NoError(t, err)
}
