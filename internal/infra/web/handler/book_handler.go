package handler

import (
	"encoding/json"
	"net/http"
	"github.com/rs/zerolog/log"
	"github.com/JeanKasoki/library-api/internal/usecase"
)

// Ele sabe traduzir HTTP (JSON) para Go e vice-versa.
type BookHandler struct {
	API *usecase.CreateBookUseCase
}

func NewBookHandler(api *usecase.CreateBookUseCase) *BookHandler{
	return &BookHandler{API: api}
}

func (bookHandler *BookHandler) Create(w http.ResponseWriter, req *http.Request) {
	log.Debug().Msg("A requisição chegou no controller")
	
	// passo 1: criar DTO do book
	var inputBook usecase.BookInput

	// passo 2: Criamos um decodificador que olha para o corpo da requisição (req.Body)
	// E tentamos converter o JSON para a struct inputBook
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&inputBook)
	if err != nil{
		// LOG: Error. O cliente mandou um JSON quebrado
		log.Error().Msg("Application have been failed")
		// Resposta HTTP 400 (Culpa do cliente)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// passo 3: Passamos os dados limpos pro UseCase validar e mandar salvar
	err = bookHandler.API.ExecuteAsync(&inputBook)
	if err != nil{
		// LOG: Error. Ocorreu erro de regra de negócio
		log.Error().Msg("Application have been failed")
		// Resposta HTTP 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}
// Se chegou aqui, o livro foi criado e salvo no banco.
	log.Info().Msgf("Livro '%s' criado com sucesso!", inputBook.Titulo)

	// passo 4: responder o cliente, Resposta HTTP 201 (Created)
	w.WriteHeader(http.StatusCreated)
}