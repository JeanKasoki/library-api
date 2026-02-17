package usecase

import "github.com/JeanKasoki/library-api/internal/infra/repository"

type DeleteBookUseCase struct {
	Repo *repository.BookRepository
}

func NewDeleteBookUseCase(repo *repository.BookRepository) *DeleteBookUseCase{
	return &DeleteBookUseCase{Repo: repo}
}

func (uc *DeleteBookUseCase) Execute(id int) error{
	err := uc.Repo.Delete(id)
	if err != nil{
		return err
	}
	return nil
}