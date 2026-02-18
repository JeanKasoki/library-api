package main

import (
	"net/http"

	"github.com/JeanKasoki/library-api/internal/entity"
	"github.com/JeanKasoki/library-api/internal/infra/database"
	"github.com/JeanKasoki/library-api/internal/infra/repository"
	"github.com/JeanKasoki/library-api/internal/infra/web/handler"
	"github.com/JeanKasoki/library-api/internal/usecase"
	"github.com/rs/zerolog/log"
)

func main(){
	// Capturamos a conexão (db) e o erro (err) para tentar conectar no banco
	db, err := database.ConnectDB()
	if err != nil{
		log.Fatal().Err(err).Msg("Erro ao conectar no banco")
	}
	log.Info().Msg("Application have been successfully")
	// Criação das auto migrations, passamos o db de cima capturado e passamos um PONTEIRO (&) de cada entidade vazia
	err = db.AutoMigrate(&entity.Book{}, &entity.Loan{}, &entity.Notification{}, &entity.User{})
	if err != nil{
		log.Fatal().Err(err).Msg("Falha ao criar tabelas (AutoMigrate)")
	}
	log.Info().Msg("Conexão e Tabelas OK")

	// --- INJEÇÃO DE DEPENDÊNCIA (LIGANDO OS CABOS) ---
 //  --- CREATE --- 
	// Criamos o Repository e damos a chave do banco (db) pra ele
	bookRepo := repository.NewBookRepository(db)
// Criamos o UseCase de Criação e apresentamos o repository (bookRepo) pra ele
	createBookUseCase := usecase.NewCreateBookUseCase(bookRepo)

	// Criamos o UseCase de Listagem e apresentamos o repository (bookRepo) pra ele
	listBooksUseCase := usecase.NewListBooksUseCase(bookRepo)

	// Criamos o UseCase de Busca e apresentamos o repository (bookRepo) pra ele
	getBookUseCase := usecase.NewGetBookUseCase(bookRepo)

	// Criamos o UseCase de Atualização e apresentamos o repository (bookRepo) pra ele
	updateBookUseCase := usecase.NewUpdateBooksUseCase(bookRepo)

	// Criamos o UseCase de Exclusão e apresentamos o repository (bookRepo) pra ele
	deleteBookUseCase := usecase.NewDeleteBookUseCase(bookRepo)

	// Criamos o Handler (o porteiro) e entregamos TODOS os 5 UseCases (as ferramentas) na mão dele
	bookHandler := handler.NewBookHandler(createBookUseCase, listBooksUseCase, getBookUseCase, updateBookUseCase, deleteBookUseCase)

	// Configuração da rota
	// "Quando alguém chamar POST /books, passa a ligação para o Handler"
	http.HandleFunc("POST /books", bookHandler.CreateBook)
	// "Quando alguém chamar GET /books, passa a ligação para o Handler listar todos os livros"
	http.HandleFunc("GET /books", bookHandler.ListBooks)
	// "Quando alguém chamar GET /book (com ?id=...), passa a ligação para o Handler buscar um livro específico"
	http.HandleFunc("GET /book", bookHandler.GetBook)
	// "Quando alguém chamar PUT /book (com ?id=...), passa a ligação para o Handler atualizar os dados do livro"
	http.HandleFunc("PUT /book", bookHandler.UpdateBook)
// "Quando alguém chamar DELETE /book (com ?id=...), passa a ligação para o Handler apagar o livro do banco"
	http.HandleFunc("DELETE /book", bookHandler.DeleteBook)

	log.Info().Msg("Servidor rodando na porta 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Error().Msg("Erro ao subir servidor HTTP")
	}
}