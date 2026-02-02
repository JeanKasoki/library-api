package usecase

import(
	"github.com/JeanKasoki/library-api/internal/infra/repository"
)

type BookOutputDTO struct {
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Autor string	`json:"autor"`
	ISBN string	`json:"isbn"`
	AnoPublicacao int	`json:"ano_publicacao"`
}

type GetBookUseCase struct{
	Repo *repository.BookRepository
}

// NewGetBookUseCase é a Factory (Construtor).
// Recebe o repositório pronto e retorna o UseCase instanciado.
func NewGetBookUseCase (repo *repository.BookRepository) *GetBookUseCase{
	return &GetBookUseCase{Repo: repo}
} 

// 'uc' = usecase
func (uc *GetBookUseCase) Execute(id int) (*BookOutputDTO, error){
	bookEntity, err := uc.Repo.FindByID(id)
	if err != nil{
		return nil, err
	}

	// Criamos a struct de saída já pegando o endereço de memória (&).
	// Isso faz com que a variável 'dto' nasça como um ponteiro (*BookOutputDTO).
		dto := &BookOutputDTO{
			ID: bookEntity.ID,
			Titulo: bookEntity.Titulo,
			Autor: bookEntity.Autor,
			ISBN: bookEntity.ISBN,
			AnoPublicacao: bookEntity.AnoPublicacao,
		}
	return dto, nil
}