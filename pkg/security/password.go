package security

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/rs/zerolog/log"
)

// HashPassword converte uma senha em texto puro para um hash irreversível
func HashPassword(password string) (string, error){
	sb, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}

	return string(sb), nil
}

// CheckPassword compara o hash salvo no banco com a senha digitada pelo usuário
func CheckPassword(hashDB string, passwordUser string) error{
	err := bcrypt.CompareHashAndPassword([]byte(hashDB), []byte(passwordUser))
	if err != nil{
		log.Error().Msg("Invalid email or password")
		return err
	}
	return nil
}