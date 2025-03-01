package models

type ExercicioTreino struct {
	ID                   uint64 `json:"id"`
	QuantidadeRepeticoes int8   `json:"quantidadeRepeticoes"`
	Minimo               int8   `json:"minimo"`
	Maximo               int8   `json:"maximo"`
	IdTreino             uint64 `json:"idTreino"`
}
