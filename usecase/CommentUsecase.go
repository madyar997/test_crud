package usecase

import "test_crud/models"

type commentUsecase struct {
	commentRepo models.CommentRepository
}

func NewCommentUsecase(commentRepository models.CommentRepository) models.CommentUsecase {
	return &commentUsecase{
		commentRepo: commentRepository,
	}
}

func (cu commentUsecase) Get() ([]models.Comment, error) {
	res, err := cu.commentRepo.Get()
	return res, err
}

func (cu commentUsecase) Create(comment *models.Comment) error {
	err := cu.commentRepo.Create(comment)
	if err != nil {
		return err
	}
	return nil
}

func (cu commentUsecase) Update(comment *models.Comment) error {
	err := cu.commentRepo.Update(comment)
	if err != nil {
		return err
	}
	return nil
}

func (cu commentUsecase) Delete(id int64) error {
	err := cu.commentRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
