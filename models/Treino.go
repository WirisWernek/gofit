package models

type Treino struct {
	ID            int64  `json:"id"`
	Quando        string `json:"quando"`
	IdPlanoTreino int64  `json:"idPlanoTreino"`
}
