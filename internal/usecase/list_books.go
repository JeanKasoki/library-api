package usecase

import "github.com/JeanKasoki/library-api/internal/infra/repository"

// Definimos o que queremos mostrar para o mundo (JSON). DTO de saída
type BookOutput struct {
	ID	int	`json:"id"`
	Titulo string `json:"titulo"`
	Autor string	`json:"autor"`
	ISBN string	`json:"isbn"`
	AnoPublicacao int	`json:"ano_publicacao"`
}

// Struct do usecase
type ListBooksUseCase struct {
	Repo *repository.BookRepository
}

// Factory
func NewListBooksUseCase(repo *repository.BookRepository) *ListBooksUseCase{
	return &ListBooksUseCase{Repo: repo}
}

// 'u' de usecase
func (u *ListBooksUseCase) Execute() ([]BookOutput, error) {
	bookEntity, err := u.Repo.FindAll()
	if err != nil{
		return nil, err
	}

	var booksOutput []BookOutput

	for _, book := range bookEntity{
		dto := BookOutput{
			ID:	book.ID,
			Titulo: book.Titulo,
			Autor: book.Autor,
			ISBN: book.ISBN,
			AnoPublicacao: book.AnoPublicacao,
		}
		booksOutput = append(booksOutput, dto)
	}
	return booksOutput, nil
}