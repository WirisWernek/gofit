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

func GetAllPlanosTreino(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planosTreino, erro := planoTreinoRepository.GetAllPlanosTreino()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, planosTreino)
}

func GetPlanoTreinoByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

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

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planoTreino, erro := planoTreinoRepository.GetPlanoTreinoByID(ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
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

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	planoTreino.ID, erro = planoTreinoRepository.InsertPlanoTreino(planoTreino)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
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
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var planoTreino models.PlanoTreino
	erro = json.Unmarshal(requestBody, &planoTreino)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	if erro = planoTreinoRepository.UpdatePlanoTreino(ID, planoTreino); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
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

	db, erro := database.Conectar()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	planoTreinoRepository := repositories.NewRepositoryPlanoTreino(db)
	if erro = planoTreinoRepository.DeletePlanoTreinoByID(ID); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
