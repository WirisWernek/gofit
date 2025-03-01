package main

import (
	"fmt"
	"gofit/controllers/equipamento"
	"gofit/controllers/exercicio"
	"gofit/controllers/planoTreino"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	loadRoutesEquipamento(router)
	loadRoutesPlanoTreino(router)
	loadRoutesExercicio(router)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func loadRoutesEquipamento(router *mux.Router) {
	endpoint := "/equipamento"

	equipamentoController := equipamento.NewEquipamento()

	router.HandleFunc(endpoint, equipamentoController.GetAll).Methods(http.MethodGet)
	router.HandleFunc(endpoint+"/{id}", equipamentoController.GetById).Methods(http.MethodGet)
	router.HandleFunc(endpoint, equipamentoController.Insert).Methods(http.MethodPost)
	router.HandleFunc(endpoint+"/{id}", equipamentoController.Update).Methods(http.MethodPut)
	router.HandleFunc(endpoint+"/{id}", equipamentoController.Delete).Methods(http.MethodDelete)

	fmt.Println("Escutando endpoints de Equipamento em /equipamento")

}

func loadRoutesPlanoTreino(router *mux.Router) {
	endpoint := "/plano-treino"

	planoTreinoController := planoTreino.NewPlanoTreino()

	router.HandleFunc(endpoint, planoTreinoController.GetAll).Methods(http.MethodGet)
	router.HandleFunc(endpoint+"/{id}", planoTreinoController.GetById).Methods(http.MethodGet)
	router.HandleFunc(endpoint, planoTreinoController.Insert).Methods(http.MethodPost)
	router.HandleFunc(endpoint+"/{id}", planoTreinoController.Update).Methods(http.MethodPut)
	router.HandleFunc(endpoint+"/{id}", planoTreinoController.Delete).Methods(http.MethodDelete)

	fmt.Println("Escutando endpoints de Plano de Treino em /plano-treino")

}

func loadRoutesExercicio(router *mux.Router) {
	endpoint := "/exercicio"

	exercicioController := exercicio.NewExercicio()

	router.HandleFunc(endpoint, exercicioController.GetAll).Methods(http.MethodGet)
	router.HandleFunc(endpoint+"/{id}", exercicioController.GetById).Methods(http.MethodGet)
	router.HandleFunc(endpoint, exercicioController.Insert).Methods(http.MethodPost)
	router.HandleFunc(endpoint+"/{id}", exercicioController.Update).Methods(http.MethodPut)
	router.HandleFunc(endpoint+"/{id}", exercicioController.Delete).Methods(http.MethodDelete)

	fmt.Println("Escutando endpoints de Exercicio em /exercicio")

}
