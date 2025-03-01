package router

import (
	"gofit/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai retornar o router com as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return routes.Configurar(router)
}
