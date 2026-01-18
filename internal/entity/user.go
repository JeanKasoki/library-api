package entity

import "time"

type User struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"-"`
	Perfil string  `json:"perfil"`
	CreatedAt time.Time `json:"created_at"`
}