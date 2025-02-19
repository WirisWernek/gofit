package models

import "time"

type PlanoTreino struct {
	ID        int64     `json:"id"`
	Inicio    time.Time `json:"inicio"`
	Ativo     bool      `json:"ativo"`
	Nome      string    `json:"nome"`
	Descricao string    `json:"descricao"`
	Descanso  int64     `json:"descanso"`
}
