package configuration

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexao é a string de conexão com o postgres
	StringConexao = ""

	// Porta é a porta onde a API vai estar rodando
	Porta = 0

	// SecretKey é a chave que será utilizada como assinatura do token
	SecretKey []byte
)

// Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	// SecretKey = []byte(os.Getenv("SECRET_KEY"))

	if erro != nil {
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORTA"),
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
