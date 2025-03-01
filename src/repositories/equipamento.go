package repositories

import (
	"database/sql"
	"gofit/models"
)

type EquipamentoRepository struct {
	db *sql.DB
}

func NewRepositoryEquipamento(db *sql.DB) *EquipamentoRepository {
	return &EquipamentoRepository{db}
}

func (repository EquipamentoRepository) InsertEquipamento(equipamento models.Equipamento) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO equipamento (nome, tipo) VALUES($1, $2) RETURNING id")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	lastID := 0
	erro = statement.QueryRow(equipamento.Nome, equipamento.Tipo).Scan(&lastID)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository EquipamentoRepository) GetAllEquipamentos() ([]models.Equipamento, error) {
	linhas, erro := repository.db.Query("SELECT * FROM equipamento")

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var equipamentos []models.Equipamento

	for linhas.Next() {
		var equipamento models.Equipamento

		if erro := linhas.Scan(&equipamento.ID, &equipamento.Nome, &equipamento.Tipo); erro != nil {
			return nil, erro
		}

		equipamentos = append(equipamentos, equipamento)
	}

	return equipamentos, nil
}

func (repository EquipamentoRepository) GetEquipamentoByID(equipamentoID uint64) (models.Equipamento, error) {
	statement, erro := repository.db.Prepare("SELECT * FROM equipamento WHERE id = $1")

	if erro != nil {
		return models.Equipamento{}, erro
	}

	defer statement.Close()

	linhas, erro := statement.Query(equipamentoID)

	if erro != nil {
		return models.Equipamento{}, erro
	}

	defer linhas.Close()

	var equipamento models.Equipamento
	if linhas.Next() {
		if erro = linhas.Scan(&equipamento.ID, &equipamento.Nome, &equipamento.Tipo); erro != nil {
			return models.Equipamento{}, erro
		}

	}

	return equipamento, nil
}

func (repository EquipamentoRepository) UpdateEquipamento(equipamentoID uint64, equipamento models.Equipamento) error {
	statement, erro := repository.db.Prepare("UPDATE equipamento SET nome = $2, tipo = $3 WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(equipamentoID, equipamento.Nome, equipamento.Tipo)

	if erro != nil {
		return erro
	}

	return nil
}

func (repository EquipamentoRepository) DeleteEquipmentoByID(equipamentoID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM equipamento WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(equipamentoID)

	if erro != nil {
		return erro
	}

	return nil
}
