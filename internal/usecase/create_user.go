package usecase

import (
	"errors" //pacote padrão para criar erros
	"github.com/JeanKasoki/library-api/internal/entity"
	"github.com/JeanKasoki/library-api/internal/infra/repository"
	"github.com/JeanKasoki/library-api/pkg/security"
)

type UserInput struct {
	Nome   string `json:"nome"`
	Email  string `json:"email"`
	Senha  string `json:"password"` // Aqui usamos "password" para ler o JSON corretamente
	Perfil string `json:"perfil"`
}

type CreateUserUseCase struct {
	Repo *repository.UserRepository
}

func NewCreateUserUseCase(repo *repository.UserRepository) *CreateUserUseCase{
	return &CreateUserUseCase{Repo: repo}
}

func (uc *CreateUserUseCase) Execute(input UserInput) error{
	//-------- ÁREA DE VALIDAÇÃO: REGRAS DE NEGÓCIO --------//
	if input.Nome == ""{
		return errors.New("o nome não pode estar vazio")
	}
	if input.Email == ""{
		return errors.New("o e-mail não pode estar vazio")
	}
	if input.Senha == ""{
		return errors.New("a senha não pode estar vazia")
	}
	if input.Perfil == ""{
		return errors.New("o perfil não pode estar vazio")
	}
	//-------- FIM DA VALIDAÇÃO --------//

	//-------- SEGURANÇA (HASH) --------//
	// Só hasheamos a senha se os dados passaram na validação
	hash, err := security.HashPassword(input.Senha)
	if err != nil{
		return err
	}

//-------- MONTAGEM DA ENTIDADE --------//
	newUser := entity.User{
		Nome: input.Nome,
		Email: input.Email,
		Senha: hash, // Aqui entra o hash gerado, e não a senha em texto puro
		Perfil: input.Perfil,
	}
	
	return uc.Repo.Create(&newUser)
}