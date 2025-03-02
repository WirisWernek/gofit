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

func GetAllExercicioPlanoTreinos(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	exercicioPlanoTreinoRepository := repositories.NewRepositoryExercicioPlanoTreino(db)
	exercicioPlanoTreinos, erro := exercicioPlanoTreinoRepository.GetAllExercicioPlanoTreinos()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, exercicioPlanoTreinos)
}

func GetExercicioPlanoTreinoByID(w http.ResponseWriter, r *http.Request) {
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

	exercicioPlanoTreinoRepository := repositories.NewRepositoryExercicioPlanoTreino(db)
	exercicio, erro := exercicioPlanoTreinoRepository.GetExercicioPlanoTreinoByID(ID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, exercicio)

}

func InsertExercicioPlanoTreino(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var exercicioPlanoTreino models.ExercicioPlanoTreino
	if erro = json.Unmarshal(requestBody, &exercicioPlanoTreino); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	exercicioPlanoTreinoRepository := repositories.NewRepositoryExercicioPlanoTreino(db)
	exercicioPlanoTreino.ID, erro = exercicioPlanoTreinoRepository.InsertExercicioPlanoTreino(exercicioPlanoTreino)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, exercicioPlanoTreino)

}

func UpdateExercicioPlanoTreino(w http.ResponseWriter, r *http.Request) {
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

	var exercicioPlanoTreino models.ExercicioPlanoTreino
	if erro = json.Unmarshal(requestBody, &exercicioPlanoTreino); erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()
	exercicioPlanoTreinoRepository := repositories.NewRepositoryExercicioPlanoTreino(db)
	if erro = exercicioPlanoTreinoRepository.UpdateExercicioPlanoTreino(ID, exercicioPlanoTreino); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func DeleteExercicioPlanoTreinoByID(w http.ResponseWriter, r *http.Request) {
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

	exercicioPlanoTreinoRepository := repositories.NewRepositoryExercicioPlanoTreino(db)
	if erro = exercicioPlanoTreinoRepository.DeleteExercicioPlanoTreinoByID(ID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
