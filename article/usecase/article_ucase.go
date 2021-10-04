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

func (a articleUsecase) GET() []models.Article {
	res := a.articleRepo.Get()
	return res
}