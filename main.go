package main

import (
	"fmt"

	"gofit/src/configuration"
	"gofit/src/router"
	"net/http"
)

func main() {

	configuration.Carregar()

	router := router.Gerar()

	fmt.Println(fmt.Sprintf("Escutando na porta %d", configuration.Porta))

	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Porta), router)
}
