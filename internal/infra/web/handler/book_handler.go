package handler

import (
	"encoding/json"
	"net/http"
	"github.com/rs/zerolog/log"
	"github.com/JeanKasoki/library-api/internal/usecase"
)

// Ele sabe traduzir HTTP (JSON) para Go e vice-versa.
type BookHandler struct {
	CreateBookUseCase *usecase.CreateBookUseCase
	ListBooksUseCase *usecase.ListBooksUseCase
}

func NewBookHandler(create *usecase.CreateBookUseCase, list *usecase.ListBooksUseCase) *BookHandler{
	return &BookHandler{
		CreateBookUseCase: create,
		ListBooksUseCase: list,
	}
}

func (h *BookHandler) Create(w http.ResponseWriter, req *http.Request) {
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
	err = h.CreateBookUseCase.Execute(&inputBook)
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

func (h *BookHandler) List (w http.ResponseWriter, req *http.Request){
	// 1. Chamar o UseCase (h.ListBooksUseCase.Execute())
	list, err := h.ListBooksUseCase.Execute()
    if err != nil {
			log.Error().Msg("Falha ao encontrar livros")
			// Respondendo 500 para o usuário
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// // Log de sucesso (apenas informativo). Usamos len(list) para dizer quantos livros achamos
		log.Info().Msgf("Sucesso! Livros encontrados: %d", len(list))
    // 3. Sucesso (Responder JSON)
		// A: Avisar o navegador/Postman que o que vem aí é JSON
		w.Header().Set("Content-Type", "application/json")
		// B: Definir status 200 (OK)
		w.WriteHeader(http.StatusOK)
    // C: Converter a lista (Go) para JSON (Texto) e escrever na resposta (w)
		// NewEncoder conecta com a saída (w). Encode pega o dado (list).
		json.NewEncoder(w).Encode(list)
}