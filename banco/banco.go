package banco

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver de Conexão com o mysql
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "host=localhost port=5432 user=postgres password=postgres dbname=gofit sslmode=disable"

	db, erro := sql.Open("postgres", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
