package routes

import (
	"gofit/src/controllers"
	"net/http"
)

var rotasExercicioPlanoTreino = []Rota{
	{
		URI:                "/exercicio-plano-treino",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InsertExercicioPlanoTreino,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio-plano-treino",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetAllExercicioPlanoTreinos,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio-plano-treino/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetExercicioPlanoTreinoByID,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio-plano-treino/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateExercicioPlanoTreino,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio-plano-treino/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteExercicioPlanoTreinoByID,
		RequerAutenticacao: false,
	},
}
