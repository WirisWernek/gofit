package models

type Exercicio struct {
	ID                         int64  `json:"id"`
	Nome                       string `json:"nome"`
	QuantidadeMinimaRepeticoes int8   `json:"quantidadeMinimaRepeticoes"`
	QuantidadeMaximaRepeticoes int8   `json:"quantidadeMaximaRepeticoes"`
	IdEquipamento              int64  `json:"idEquipamento"`
	IdPlanoTreino              int64  `json:"idPlanoTreino"`
}
