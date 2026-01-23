package usecase

import (
	"github.com/JeanKasoki/library-api/internal/entity"
	"github.com/JeanKasoki/library-api/internal/infra/repository"
)

// Dados que o usuário vai mandar
type BookInput struct{
	Titulo string
	Autor string
	ISBN string
	AnoPublicacao int
}

// Estrutura do caso de uso da criação de livro
type CreateBookUseCase struct {
	Repo *repository.BookRepository
}

func NewCreateBookUseCase(repo *repository.BookRepository) *CreateBookUseCase{
	return &CreateBookUseCase{Repo: repo}
}

func (u *CreateBookUseCase) Execute(input *BookInput) error{
	newBook := entity.Book{
		Titulo: input.Titulo,
		Autor: input.Autor,
		ISBN: input.ISBN,
		AnoPublicacao: input.AnoPublicacao,
	}
		// O 'u' é o UseCase. Dentro dele tem o 'Repo'. Dentro do Repo tem o 'Create'.
    // O Create espera um PONTEIRO de livro (&newBook).
	err := u.Repo.Create(&newBook)
	return err
}