package models

type ExercicioTreino struct {
	ID                   int64 `json:"id"`
	QuantidadeRepeticoes int8  `json:"quantidadeRepeticoes"`
	Minimo               int8  `json:"minimo"`
	Maximo               int8  `json:"maximo"`
	IdTreino             int64 `json:"idTreino"`
}
