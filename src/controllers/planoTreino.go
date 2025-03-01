package controllers

import (
	"encoding/json"
	"gofit/banco"
	"gofit/models"
	"gofit/src/repositories"
	"gofit/src/response"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPlanosTreino(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planosTreino, erro := planoTreinoRepository.GetAllPlanosTreino()
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusCreated, planosTreino)
}

func GetPlanoTreinoByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planoTreino, erro := planoTreinoRepository.GetPlanoTreinoByID(ID)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, planoTreino)
}

func InsertPlanoTreino(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var planoTreino models.PlanoTreino
	erro = json.Unmarshal(requestBody, &planoTreino)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planoTreino.ID, erro = planoTreinoRepository.InsertPlanoTreino(planoTreino)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusCreated, planoTreino)
}

func UpdatePlanoTreino(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var planoTreino models.PlanoTreino
	erro = json.Unmarshal(requestBody, &planoTreino)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	if erro = planoTreinoRepository.UpdatePlanoTreino(ID, planoTreino); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeletePlanoTreinoByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	if erro = planoTreinoRepository.DeletePlanoTreinoByID(ID); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
