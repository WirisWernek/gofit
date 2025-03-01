package routes

import (
	"gofit/src/controllers"
	"net/http"
)

var rotasEquipamento = []Rota{
	{
		URI:                "/equipamento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.InsertEquipamento,
		RequerAutenticacao: false,
	},
	{
		URI:                "/equipamento",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetAllEquipamentos,
		RequerAutenticacao: false,
	},
	{
		URI:                "/equipamento/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetEquipamentoByID,
		RequerAutenticacao: false,
	},
	{
		URI:                "/equipamento/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateEquipamento,
		RequerAutenticacao: false,
	},
	{
		URI:                "/equipamento/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteEquipmentoByID,
		RequerAutenticacao: false,
	},
}
