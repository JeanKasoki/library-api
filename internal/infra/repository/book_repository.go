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

// FindByID busca um livro específico pelo ID
func (repo *BookRepository) FindByID(id int) (*entity.Book, error){
	var book entity.Book

	err := repo.DB.First(&book, id).Error
	if err != nil{
		return nil, err
	}
	return &book, nil
}

// Update atualiza os dados de um livro no banco.
func (repo *BookRepository) Update(book *entity.Book) (*entity.Book, error){
	// 1. Usamos Save() porque ele lida com "Update se existir ID, Create se não existir"
  // 2. Passamos 'book' direto, sem &, pois ele já é um ponteiro
	err := repo.DB.Save(book).Error
	if err != nil{
		return nil, err
	}
	return book, nil
}

// Delete remove um livro do banco de dados a partir do seu ID
func (repo *BookRepository) Delete(id int) error{
	err := repo.DB.Delete(&entity.Book{}, id).Error
	if err != nil{
		return err
	}
	return nil
}