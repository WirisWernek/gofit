package models

type Equipamento struct {
	ID   uint64 `json:"id"`
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
}
