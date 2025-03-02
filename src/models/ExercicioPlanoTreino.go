package models

type ExercicioPlanoTreino struct {
	ID                         uint64 `json:"id"`
	ExercicioID                uint64 `json:"idExercicio"`
	PlanoTreinoID              uint64 `json:"idPlanoTreino"`
	QuantidadeMinimaRepeticoes int8   `json:"quantidadeMinimaRepeticoes"`
	QuantidadeMaximaRepeticoes int8   `json:"quantidadeMaximaRepeticoes"`
}
