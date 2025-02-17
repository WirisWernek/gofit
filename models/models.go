package models

type Equipamento struct {
	ID   int64  `json:"id"`
	Nome string `json:"nome"`
	Tipo string `json:"tipo"`
}

type PlanoTreino struct {
	ID     int64  `json:"id"`
	Inicio string `json:"inicio"`
	Ativo  bool   `json:"ativo"`
}

type Treino struct {
	ID            int64  `json:"id"`
	Quando        string `json:"quando"`
	IdPlanoTreino int64  `json:"idPlanoTreino"`
}

type Exercicio struct {
	ID                         int64  `json:"id"`
	Nome                       string `json:"nome"`
	QuantidadeMinimaRepeticoes int8   `json:"quantidadeMinimaRepeticoes"`
	QuantidadeMaximaRepeticoes int8   `json:"quantidadeMaximaRepeticoes"`
	IdEquipamento              int64  `json:"idEquipamento"`
	IdPlanoTreino              int64  `json:"idPlanoTreino"`
}

type ExercicioTreino struct {
	ID                   int64 `json:"id"`
	QuantidadeRepeticoes int8  `json:"quantidadeRepeticoes"`
	Minimo               int8  `json:"minimo"`
	Maximo               int8  `json:"maximo"`
	IdTreino             int64 `json:"idTreino"`
}
