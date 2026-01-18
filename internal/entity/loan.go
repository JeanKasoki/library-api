package entity

import(
	"time"
)

type Loan struct{
	ID int `json:"id"`
	UsuarioID int `json:"id_usuario"`
	LivroID int `json:"id_livro"`
	DataReserva time.Time `json:"data_reserva"`
	DataRetirada *time.Time `json:"data_retirada"`
	DataDevolucaoPrevista *time.Time `json:"data_devolucao_prevista"`
	DataDevolucaoReal *time.Time `json:"data_devolucao_real"`
	ValorMulta float64 `json:"valor_multa"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}