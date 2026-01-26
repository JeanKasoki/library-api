package repository

import (
	"github.com/JeanKasoki/library-api/internal/entity"
	"gorm.io/gorm"
)

// BookRepository é quem manda no banco para assuntos de Livro
type BookRepository struct {
	DB *gorm.DB
}

// NewBookRepository cria uma nova instância (fábrica)
func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// Create recebe um Livro (da pasta entity) e salva no banco
func (repo *BookRepository) Create(book *entity.Book) error {
		return repo.DB.Create(book).Error
}
// FindAll busca todos os registros de livros na tabela.
func (repo *BookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book

	// O método .Find() gera o SQL "SELECT * FROM books".
	err := repo.DB.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}