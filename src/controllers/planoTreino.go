package controllers

import (
	"encoding/json"
	"gofit/banco"
	"gofit/models"
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

	linhas, erro := db.Query("SELECT * FROM plano_treino")

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer linhas.Close()

	var planoTreinos []models.PlanoTreino

	for linhas.Next() {
		var planoTreino models.PlanoTreino

		erro = linhas.Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo)

		if erro != nil {
			response.Erro(w, http.StatusInternalServerError, erro)
			return
		}

		planoTreinos = append(planoTreinos, planoTreino)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	erro = json.NewEncoder(w).Encode(planoTreinos)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
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

	var planoTreino models.PlanoTreino
	erro = db.QueryRow("SELECT * FROM plano_treino WHERE id = $1", ID).Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	erro = json.NewEncoder(w).Encode(planoTreino)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
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

	statement, erro := db.Prepare("INSERT INTO plano_treino (nome, descricao, inicio, descanso, ativo) VALUES($1, $2, $3, $4, $5)")

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	_, erro = insercao.RowsAffected()

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	message, erro := json.Marshal("Equipamento inserido")

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(message)
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

	statement, erro := db.Prepare("UPDATE plano_treino SET nome=$2, descricao=$3, inicio=$4, descanso=$5, ativo=$6 WHERE id=$1")

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(ID, planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	statement, erro := db.Prepare("DELETE FROM plano_treino WHERE id = $1")

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(ID)

	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
