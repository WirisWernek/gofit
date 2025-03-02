package models

type Exercicio struct {
	ID            uint64 `json:"id"`
	Nome          string `json:"nome"`
	EquipamentoID uint64 `json:"idEquipamento"`
}
