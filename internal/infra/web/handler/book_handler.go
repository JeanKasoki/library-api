package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/rs/zerolog/log"
	"github.com/JeanKasoki/library-api/internal/usecase"
)

// Ele sabe traduzir HTTP (JSON) para Go e vice-versa.
type BookHandler struct {
	CreateBookUseCase *usecase.CreateBookUseCase
	ListBooksUseCase *usecase.ListBooksUseCase
	GetBookUseCase *usecase.GetBookUseCase
	UpdateBookUseCase *usecase.UpdateBookUseCase
	DeleteBookUseCase *usecase.DeleteBookUseCase
}

func NewBookHandler(create *usecase.CreateBookUseCase, list *usecase.ListBooksUseCase, get *usecase.GetBookUseCase, update *usecase.UpdateBookUseCase, delete *usecase.DeleteBookUseCase) *BookHandler{
	return &BookHandler{
		CreateBookUseCase: create,
		ListBooksUseCase: list,
		GetBookUseCase: get,
		UpdateBookUseCase: update,
		DeleteBookUseCase: delete,
	}
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, req *http.Request) {
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



func (h *BookHandler) ListBooks(w http.ResponseWriter, req *http.Request){
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



// Novo Método: GET /books/find?id=1
func (h *BookHandler) GetBook(w http.ResponseWriter, req *http.Request) {
	// 1. Pegar o ID da URL (Query Parameter)
	// O req.URL.Query().Get("id") busca o valor de ?id=...
	idString := req.URL.Query().Get("id")

	// 2. Converter String para Int
	// Atoi significa (Ascii to Integer). Retorna o numero e um erro.
	id, err := strconv.Atoi(idString)
	if err != nil {
		// Se der erro (ex: id vazio ou "abc"), retornamos erro 400 (Bad Request)
		log.Error().Msg("ID inválido ou não fornecido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. Chamar o UseCase passando o ID convertido
	book, err := h.GetBookUseCase.Execute(id)
	if err != nil {
		// Aqui assumimos 500, mas idealmente seria 404 (Not Found) se não achar
		log.Error().Err(err).Msg("Falha ao buscar livro")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 4. Sucesso
	log.Info().Msgf("Livro encontrado: %s", book.Titulo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	// Correção: Passando (w) para o Encoder
	json.NewEncoder(w).Encode(book)
}



func (h *BookHandler) UpdateBook(w http.ResponseWriter, req *http.Request){
	// 1. Preparar a variável para receber o JSON
	var inputBookDTO usecase.UpdateBookInputDTO

	// 2. Pegar o ID da URL (Igual ao GetBook)
	idString := req.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil{
		log.Error().Msg("ID inválido ou não fornecido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. Ler o JSON do Corpo (Igual ao CreateBook)
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&inputBookDTO)
	if err != nil{
		// LOG: Error. O cliente mandou um JSON quebrado
		log.Error().Msg("Erro ao ler JSON: " + err.Error())
		// Resposta HTTP 400 (Culpa do cliente)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 4. Chamar o Gerente (UseCase)
	// Passamos o ID (quem alterar) e o Input (o que alterar)
	// Recebemos o Livro Atualizado (output) ou erro
	output, err := h.UpdateBookUseCase.Execute(id, inputBookDTO)
	if err != nil {
	// Se der erro aqui, pode ser que o livro não exista ou erro de banco
		log.Error().Err(err).Msg("Erro ao atualizar livro")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Sucesso
	log.Info().Msgf("Livro atualizado com sucesso: %s", output.Titulo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK (Não use 201 Created para updates)
	// Devolvemos o JSON do livro atualizado para o usuário ver
	json.NewEncoder(w).Encode(output)
}



// DeleteBook deleta um livro existente baseado no ID passado na URL
func (h *BookHandler) DeleteBook(w http.ResponseWriter, req *http.Request){
	// 1. Pegar o ID da URL (ex: /book?id=1)
	idString := req.URL.Query().Get("id")

	// 2. Tentar converter o ID de Texto (string) para Número (int)
	id, err := strconv.Atoi(idString)
	if err != nil {
		// Se der erro (ex: id vazio ou "abc"), retornamos erro 400 (Bad Request)
		log.Error().Msg("ID inválido ou não fornecido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. Chamar o Gerente (UseCase) para fazer o trabalho sujo
	err = h.DeleteBookUseCase.Execute(id)
	if err != nil {
		// Aqui assumimos 500, mas idealmente seria 404 (Not Found) se não achar
		log.Error().Err(err).Msg("Falha ao deletar livro")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Sucesso
	// Deletar não tem corpo de resposta, então NÃO setamos Content-Type
	// Apenas avisamos o terminal/log que deu certo
	log.Info().Msgf("Livro com ID %d deletado com sucesso!", id)
	// Retornamos o Status 204 (No Content)
	// Significa: "Ação concluída com sucesso, e não há mais nada a enviar"
	w.WriteHeader(http.StatusNoContent) 
}