package banco

import (
	"database/sql"
	"gofit/src/configuration"

	_ "github.com/lib/pq" // Driver de Conexão com o PostgreSQL
)

// Conectar abre a conexao com o banco de dados
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("postgres", configuration.StringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
