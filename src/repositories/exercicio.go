package repositories

import (
	"database/sql"
	"gofit/src/models"
)

type ExercicioRepository struct {
	db *sql.DB
}

func NewRepositoryExercicio(db *sql.DB) *ExercicioRepository {
	return &ExercicioRepository{db}
}

func (repository ExercicioRepository) InsertExercicio(exercicio models.Exercicio) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO exercicio (nome, quantidade_minima_repeticoes, quantidade_maxima_repeticoes, id_equipamento) VALUES($1, $2, $3, $4) RETURNING id")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	lastID := 0
	erro = statement.QueryRow(exercicio.Nome, exercicio.QuantidadeMinimaRepeticoes, exercicio.QuantidadeMaximaRepeticoes, exercicio.EquipamentoID).Scan(&lastID)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository ExercicioRepository) GetAllExercicios() ([]models.Exercicio, error) {
	linhas, erro := repository.db.Query("SELECT * FROM exercicio")

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var exercicios []models.Exercicio

	for linhas.Next() {
		var exercicio models.Exercicio

		if erro := linhas.Scan(&exercicio.ID, &exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID); erro != nil {
			return nil, erro
		}

		exercicios = append(exercicios, exercicio)
	}

	return exercicios, nil
}

func (repository ExercicioRepository) GetExercicioByID(exercicioID uint64) (models.Exercicio, error) {
	statement, erro := repository.db.Prepare("SELECT * FROM exercicio WHERE id = $1")

	if erro != nil {
		return models.Exercicio{}, erro
	}

	defer statement.Close()

	linhas, erro := statement.Query(exercicioID)

	if erro != nil {
		return models.Exercicio{}, erro
	}

	defer linhas.Close()

	var exercicio models.Exercicio
	if linhas.Next() {
		if erro = linhas.Scan(&exercicio.ID, &exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID); erro != nil {
			return models.Exercicio{}, erro
		}

	}

	return exercicio, nil
}

func (repository ExercicioRepository) UpdateExercicio(exercicioID uint64, exercicio models.Exercicio) error {
	statement, erro := repository.db.Prepare("UPDATE exercicio SET nome= $2, quantidade_minima_repeticoes=$3, quantidade_maxima_repeticoes=$4, id_equipamento=$5 WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(exercicioID, exercicio.Nome, exercicio.QuantidadeMinimaRepeticoes, exercicio.QuantidadeMaximaRepeticoes, exercicio.EquipamentoID)

	if erro != nil {
		return erro
	}

	return nil
}

func (repository ExercicioRepository) DeleteExercicioByID(exercicioID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM exercicio WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(exercicioID)

	if erro != nil {
		return erro
	}

	return nil
}
