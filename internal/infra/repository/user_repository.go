package repository

import (
	"github.com/JeanKasoki/library-api/internal/entity"
	"gorm.io/gorm"
)

// UserRepository guarda a conexão com o banco de dados.
// É a nossa porta de entrada para manipular a tabela de usuários.
type UserRepository struct{
	DB *gorm.DB
}

// NewUserRepository é o construtor (factory). Ele recebe a conexão aberta
// lá da main.go e injeta dentro do nosso repositório.
func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{DB:db}
}

// Create recebe um usuário (com a senha já transformada em hash pelo UseCase)
// e salva ele definitivamente no banco de dados.
func (repo *UserRepository) Create(user *entity.User) error{
	return repo.DB.Create(user).Error
}

// FindByEmail busca um usuário no banco através do e-mail.
// Retorna um ponteiro (*entity.User) para podermos devolver 'nil' caso não encontre ninguém.
func (repo *UserRepository) FindByEmail(email string) (*entity.User, error){
	// 1. Criamos uma variável vazia. O GORM vai usá-la como um "molde" para despejar os dados que ele encontrar no banco.
	var user entity.User

	// 2. Fazemos a busca: "Onde a coluna email for igual à variável email, pegue o Primeiro (First)"
	err := repo.DB.Where("email = ?", email).First(&user).Error
	if err != nil{
		return nil, err
	}
	// Retornamos o endereço de memória (&) da variável preenchida e 'nil' para o erro.
	return &user, nil
}