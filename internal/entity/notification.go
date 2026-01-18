package entity

import "time"

type Notification struct{
	ID int `json:"id"`
	UsuarioID int `json:"id_usuario"`
	Mensagem string `json:"mensagem"`
	DataEnvio time.Time `json:"data_envio"`
}