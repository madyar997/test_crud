package usecase

import "test_crud/models"

type authorUsecase struct {
	authorRepo models.AuthorRepository
}

func NewAuthorUsecase(a models.AuthorRepository) models.AuthorUsecase {
	return &authorUsecase{
		authorRepo: a,
	}
}

func (a authorUsecase) Get() ([]models.Author, error) {
	res, err := a.authorRepo.Get()
	return res, err
}

func (a authorUsecase) Create(author *models.Author) error {
	err := a.authorRepo.Create(author)
	if err != nil {
		return err
	}
	return nil
}

func (a authorUsecase) Update(author *models.Author) error {
	err := a.authorRepo.Update(author)
	if err != nil {
		return err
	}
	return nil
}

func (a authorUsecase) Delete(id int64) error {
	err := a.authorRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
