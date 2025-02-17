package main

import (
	"fmt"
	"gofit/equipamento"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	loadRoutesEquipamento(router)

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

func loadRoutesEquipamento(router *mux.Router) {
	endpoint := "/equipamento"

	router.HandleFunc(endpoint, equipamento.GetAll).Methods(http.MethodGet)
	router.HandleFunc(endpoint+"/{id}", equipamento.GetById).Methods(http.MethodGet)
	router.HandleFunc(endpoint, equipamento.Insert).Methods(http.MethodPost)
	router.HandleFunc(endpoint+"/{id}", equipamento.Update).Methods(http.MethodPut)
	router.HandleFunc(endpoint+"/{id}", equipamento.Delete).Methods(http.MethodDelete)

	fmt.Println("Escutando endpoints de Equipamento em /equipamento")

}
