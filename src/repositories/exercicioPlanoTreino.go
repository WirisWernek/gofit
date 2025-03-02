package repositories

import (
	"database/sql"
	"gofit/src/models"
)

type ExercicioPlanoTreinoRepository struct {
	db *sql.DB
}

func NewRepositoryExercicioPlanoTreino(db *sql.DB) *ExercicioPlanoTreinoRepository {
	return &ExercicioPlanoTreinoRepository{db}
}

func (repository ExercicioPlanoTreinoRepository) InsertExercicioPlanoTreino(exercicioPlanoTreino models.ExercicioPlanoTreino) (uint64, error) {
	statement, erro := repository.db.Prepare("INSERT INTO exercicio_plano_treino (id_exercicio, id_plano_treino, quantidade_minima_repeticoes, quantidade_maxima_repeticoes) VALUES($1, $2, $3, $4) RETURNING id")

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	lastID := 0
	erro = statement.QueryRow(exercicioPlanoTreino.ExercicioID, exercicioPlanoTreino.PlanoTreinoID, exercicioPlanoTreino.QuantidadeMinimaRepeticoes, exercicioPlanoTreino.QuantidadeMaximaRepeticoes).Scan(&lastID)

	if erro != nil {
		return 0, erro
	}

	return uint64(lastID), nil
}

func (repository ExercicioPlanoTreinoRepository) GetAllExercicioPlanoTreinos() ([]models.ExercicioPlanoTreino, error) {
	linhas, erro := repository.db.Query("SELECT id, id_exercicio, id_plano_treino, quantidade_minima_repeticoes, quantidade_maxima_repeticoes FROM exercicio_plano_treino")

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var exerciciosPlanoTreino []models.ExercicioPlanoTreino

	for linhas.Next() {
		var exercicioPlanoTreino models.ExercicioPlanoTreino

		if erro := linhas.Scan(&exercicioPlanoTreino.ID, &exercicioPlanoTreino.ExercicioID, &exercicioPlanoTreino.PlanoTreinoID, &exercicioPlanoTreino.QuantidadeMinimaRepeticoes, &exercicioPlanoTreino.QuantidadeMaximaRepeticoes); erro != nil {
			return nil, erro
		}

		exerciciosPlanoTreino = append(exerciciosPlanoTreino, exercicioPlanoTreino)
	}

	return exerciciosPlanoTreino, nil
}

func (repository ExercicioPlanoTreinoRepository) GetExercicioPlanoTreinoByID(exercicioPlanoTreinoID uint64) (models.ExercicioPlanoTreino, error) {
	statement, erro := repository.db.Prepare("SELECT id, id_exercicio, id_plano_treino, quantidade_minima_repeticoes, quantidade_maxima_repeticoes FROM exercicio_plano_treino WHERE id = $1")

	if erro != nil {
		return models.ExercicioPlanoTreino{}, erro
	}

	defer statement.Close()

	linhas, erro := statement.Query(exercicioPlanoTreinoID)

	if erro != nil {
		return models.ExercicioPlanoTreino{}, erro
	}

	defer linhas.Close()

	var exercicioPlanoTreino models.ExercicioPlanoTreino

	if linhas.Next() {

		if erro := linhas.Scan(&exercicioPlanoTreino.ID, &exercicioPlanoTreino.ExercicioID, &exercicioPlanoTreino.PlanoTreinoID, &exercicioPlanoTreino.QuantidadeMinimaRepeticoes, &exercicioPlanoTreino.QuantidadeMaximaRepeticoes); erro != nil {
			return models.ExercicioPlanoTreino{}, erro
		}

	}

	return exercicioPlanoTreino, nil
}

func (repository ExercicioPlanoTreinoRepository) UpdateExercicioPlanoTreino(exercicioPlanoTreinoID uint64, exercicioPlanoTreino models.ExercicioPlanoTreino) error {
	statement, erro := repository.db.Prepare("UPDATE exercicio_plano_treino SET id_exercicio= $2, id_plano_treino= $3, quantidade_minima_repeticoes= $4, quantidade_maxima_repeticoes=$5 WHERE id= $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(exercicioPlanoTreinoID, exercicioPlanoTreino.ExercicioID, exercicioPlanoTreino.PlanoTreinoID, exercicioPlanoTreino.QuantidadeMinimaRepeticoes, exercicioPlanoTreino.QuantidadeMaximaRepeticoes)

	if erro != nil {
		return erro
	}

	return nil
}

func (repository ExercicioPlanoTreinoRepository) DeleteExercicioPlanoTreinoByID(exercicioPlanoTreinoID uint64) error {
	statement, erro := repository.db.Prepare("DELETE FROM exercicio_plano_treino WHERE id = $1")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	_, erro = statement.Exec(exercicioPlanoTreinoID)

	if erro != nil {
		return erro
	}

	return nil
}
