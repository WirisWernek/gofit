package main

import (
	"fmt"
	"gofit/controllers/planoTreino"
	"gofit/equipamento"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	loadRoutesEquipamento(router)
	loadRoutesPlanoTreino(router)

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
