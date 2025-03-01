package repositories

import (
	"database/sql"
	"gofit/models"
)

type PlanoTreinoRepository struct {
	db *sql.DB
}

func NewRepositoryPlanoTreino(db *sql.DB) *PlanoTreinoRepository {
	return &PlanoTreinoRepository{db}
}

func (repository PlanoTreinoRepository) InsertPlanoTreino(planoTreino models.PlanoTreino) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO plano_treino (nome, descricao, inicio, descanso, ativo) VALUES($1, $2, $3, $4, $5) RETURNING id")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	lastID := 0
	erro = statement.QueryRow(planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo).Scan(&lastID)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository PlanoTreinoRepository) GetAllPlanosTreino() ([]models.PlanoTreino, error) {
	linhas, erro := repository.db.Query("SELECT * FROM plano_treino")

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var planoTreinos []models.PlanoTreino

	for linhas.Next() {
		var planoTreino models.PlanoTreino

		if erro := linhas.Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo); erro != nil {
			return nil, erro
		}

		planoTreinos = append(planoTreinos, planoTreino)
	}

	return planoTreinos, nil
}

func (repository PlanoTreinoRepository) GetPlanoTreinoByID(planoTreinoID uint64) (models.PlanoTreino, error) {
	statement, erro := repository.db.Prepare("SELECT * FROM plano_treino WHERE id = $1")

	if erro != nil {
		return models.PlanoTreino{}, erro
	}

	defer statement.Close()

	linhas, erro := statement.Query(planoTreinoID)

	if erro != nil {
		return models.PlanoTreino{}, erro
	}

	defer linhas.Close()

	var planoTreino models.PlanoTreino
	if linhas.Next() {
		if erro = linhas.Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo); erro != nil {
			return models.PlanoTreino{}, erro
		}

	}

	return planoTreino, nil
}

func (repository PlanoTreinoRepository) UpdatePlanoTreino(planoTreinoID uint64, planoTreino models.PlanoTreino) error {
	statement, erro := repository.db.Prepare("UPDATE plano_treino SET nome=$2, descricao=$3, inicio=$4, descanso=$5, ativo=$6 WHERE id=$1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(planoTreinoID, planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo)

	if erro != nil {
		return erro
	}

	return nil
}

func (repository PlanoTreinoRepository) DeletePlanoTreinoByID(planoTreinoID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM plano_treino WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(planoTreinoID)

	if erro != nil {
		return erro
	}

	return nil
}
