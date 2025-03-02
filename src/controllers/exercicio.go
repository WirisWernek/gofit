package controllers

import (
	"encoding/json"
	"gofit/src/database"
	"gofit/src/models"
	"gofit/src/repositories"
	"gofit/src/response"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllExercicios(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	exercicioRepository := repositories.NewRepositoryExercicio(db)
	exercicios, erro := exercicioRepository.GetAllExercicios()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, exercicios)
}

func GetExercicioByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	exercicioRepository := repositories.NewRepositoryExercicio(db)
	exercicio, erro := exercicioRepository.GetExercicioByID(ID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, exercicio)

}

func InsertExercicio(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var exercicio models.Exercicio
	if erro = json.Unmarshal(requestBody, &exercicio); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	exercicioRepository := repositories.NewRepositoryExercicio(db)
	exercicio.ID, erro = exercicioRepository.InsertExercicio(exercicio)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, exercicio)

}
func UpdateExercicio(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var exercicio models.Exercicio
	if erro = json.Unmarshal(requestBody, &exercicio); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	exercicioRepository := repositories.NewRepositoryExercicio(db)
	if erro = exercicioRepository.UpdateExercicio(ID, exercicio); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
func DeleteExercicioByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	exercicioRepository := repositories.NewRepositoryExercicio(db)
	if erro = exercicioRepository.DeleteExercicioByID(ID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
