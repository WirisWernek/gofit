package models

type Exercicio struct {
	ID                         uint64 `json:"id"`
	Nome                       string `json:"nome"`
	QuantidadeMinimaRepeticoes int8   `json:"quantidadeMinimaRepeticoes"`
	QuantidadeMaximaRepeticoes int8   `json:"quantidadeMaximaRepeticoes"`
	EquipamentoID              uint64 `json:"idEquipamento"`
}
