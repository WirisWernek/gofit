package models

type Treino struct {
	ID            uint64 `json:"id"`
	Quando        string `json:"quando"`
	IdPlanoTreino uint64 `json:"idPlanoTreino"`
}
