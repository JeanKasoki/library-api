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
func (r *BookRepository) Create(book *entity.Book) error {
		return r.DB.Create(book).Error
}