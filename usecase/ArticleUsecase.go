package usecase

import (
	"test_crud/models"
)

type articleUsecase struct {
	articleRepo models.ArticleRepository
}

func NewArticleUsecase(a models.ArticleRepository) models.ArticleUsecase {
	return &articleUsecase{
		articleRepo: a,
	}
}

func (a articleUsecase) Get() ([]models.Article, error) {
	res, err := a.articleRepo.Get()
	return res, err
}

func (a articleUsecase) Create(article *models.Article) error {
	err := a.articleRepo.Create(article)
	if err != nil {
		return err
	}
	return nil
}

func (a articleUsecase) Delete(id int64) error {
	err := a.articleRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (a articleUsecase) Update(article *models.Article) error {
	err := a.articleRepo.Update(article)
	if err != nil {
		return err
	}
	return nil
}

func (a articleUsecase) GetByAuthorId(authorId int64) ([]models.Article, error) {
	res, err := a.articleRepo.GetByAuthorId(authorId)
	return res, err
}

func (a articleUsecase) GetById(id int64) (models.Article, error) {
	res, err := a.articleRepo.GetById(id)
	return res, err
}
