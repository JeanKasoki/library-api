package main

import(
	"fmt"
	"github.com/JeanKasoki/library-api/internal/entity"
	"github.com/JeanKasoki/library-api/internal/infra/database"
)

func main(){
	// Capturamos a conexão (db) e o erro (err) para tentar conectar no banco
	db, err := database.ConnectDB()
	if err != nil{
		panic(err)
	}
	fmt.Println("Conexão com o banco de dados realizada com sucesso.")
	// Criação das auto migrations, passamos o db de cima capturado e passamos um PONTEIRO (&) de cada entidade vazia
	err = db.AutoMigrate(&entity.Book{}, &entity.Loan{}, &entity.Notification{}, &entity.User{})
	if err != nil{
		panic("Falha em realizar a migração: " + err.Error())
	}
	fmt.Println("Tabelas criadas/atualizadas com sucesso.")
}