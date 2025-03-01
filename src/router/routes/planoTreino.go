package routes

import (
	"gofit/src/controllers"
	"net/http"
)

var rotasPlanoTreino = []Rota{

	{
		URI:                "/plano-treino",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InsertPlanoTreino,
		RequerAutenticacao: false,
	},
	{
		URI:                "/plano-treino",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetAllPlanosTreino,
		RequerAutenticacao: false,
	},
	{
		URI:                "/plano-treino/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetPlanoTreinoByID,
		RequerAutenticacao: false,
	},
	{
		URI:                "/plano-treino/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdatePlanoTreino,
		RequerAutenticacao: false,
	},
	{
		URI:                "/plano-treino/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletePlanoTreinoByID,
		RequerAutenticacao: false,
	},
}
