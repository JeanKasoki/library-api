package usecase

import "github.com/JeanKasoki/library-api/internal/infra/repository"

type UpdateBookInputDTO struct{
	Titulo string `json:"titulo"`
	Autor string `json:"autor"`
	ISBN string `json:"isbn"`
	AnoPublicacao int `json:"ano_publicacao"`
}

type UpdateBookUseCase struct {
	Repo *repository.BookRepository
}

func NewUpdateBooksUseCase(repo *repository.BookRepository) *UpdateBookUseCase{
	return &UpdateBookUseCase{Repo: repo}
}

func (uc *UpdateBookUseCase) Execute(id int, input UpdateBookInputDTO) (*BookOutputDTO, error){
	bookEntity, err := uc.Repo.FindByID(id)
	if err != nil{
		return nil, err
	}
	bookEntity.Titulo = input.Titulo
	bookEntity.Autor = input.Autor
	bookEntity.ISBN = input.ISBN
	bookEntity.AnoPublicacao = input.AnoPublicacao

	_, err = uc.Repo.Update(&bookEntity)
	if err != nil{
		return nil, err
	}

	dto := &BookOutputDTO{
		ID: bookEntity.ID,
		Titulo: bookEntity.Titulo,
		Autor: bookEntity.Autor,
		ISBN: bookEntity.ISBN,
		AnoPublicacao: bookEntity.AnoPublicacao,
	}
	return dto, nil
}