package models

type Equipamento struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
}
