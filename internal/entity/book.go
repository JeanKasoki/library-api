package entity

import(
	"time"
)

type Book struct{
	ID int `json:"id"`
	Titulo string `json:"titulo"`
	Autor string `json:"autor"`
	ISBN string `json:"isbn"`
	AnoPublicacao int `json:"ano_publicacao"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}