package routes

import (
	"gofit/src/controllers"
	"net/http"
)

var rotasExercicio = []Rota{
	{
		URI:                "/exercicio",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InsertExercicio,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetAllExercicios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetExercicioByID,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateExercicio,
		RequerAutenticacao: false,
	},
	{
		URI:                "/exercicio/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteExercicioByID,
		RequerAutenticacao: false,
	},
}
