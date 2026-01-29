package usecase

import (
	"errors" // Pacote padrão para criar erros
	"time"   // Para verificar ano atual
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

// BookInput (DTO - Data Transfer Object)
// Estrutura do caso de uso da criação de livro
type CreateBookUseCase struct {
	Repo *repository.BookRepository
}

func NewCreateBookUseCase(repo *repository.BookRepository) *CreateBookUseCase{
	return &CreateBookUseCase{Repo: repo}
}

func (uc *CreateBookUseCase) Execute(input *BookInput) error {
	// --- AREA DE VALIDAÇÃO (Regras de Negócio) ---
	if input.Titulo == ""{
		return errors.New("o título do livro não pode estar vazio")
	}

  if len(input.ISBN) < 10 {
		return errors.New("ISBN inválido: deve ter pelo menos 10 caracteres")
	}

	anoAtual := time.Now().Year()
	if input.AnoPublicacao > anoAtual{
		return errors.New("o ano de publicação não pode ser no futuro")
	}
	// --- FIM DA VALIDAÇÃO ---
	newBook := entity.Book{
		Titulo: input.Titulo,
		Autor: input.Autor,
		ISBN: input.ISBN,
		AnoPublicacao: input.AnoPublicacao,
	}
		// O 'uc' é o UseCase. Dentro dele tem o 'Repo'. Dentro do Repo tem o 'Create'.
    // O Create espera um PONTEIRO de livro (&newBook).
	err := uc.Repo.Create(&newBook)
	return err
}