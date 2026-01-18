package database

import (
	"log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)
// ConnectDB tenta conectar ao banco e retorna a instância do GORM
func ConnectDB() (*gorm.DB, error) {
  // DSN (Data Source Name): A "URL" de conexão do MySQL
	// parseTime=True é OBRIGATÓRIO para funcionar com as structs que usam time.Time
	// user:password@protocol(address:port)/database?settings
  dsn := "root:root@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	// Abre a conexão usando o driver do MySQL
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar no banco de dados: ", err)
			return nil, err
	}
		return db, nil
}