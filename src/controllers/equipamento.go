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

func GetAllEquipamentos(w http.ResponseWriter, r *http.Request) {
	db, erro := database.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao database"))
		return
	}

	defer db.Close()

	equipamentoRepository := repositories.NewRepositoryEquipamento(db)
	equipamentos, erro := equipamentoRepository.GetAllEquipamentos()
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusOK, equipamentos)

}

func GetEquipamentoByID(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro ID"))
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao database"))
		return
	}

	defer db.Close()

	equipamentoRepository := repositories.NewRepositoryEquipamento(db)
	equipamento, erro := equipamentoRepository.GetEquipamentoByID(ID)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusCreated, equipamento)

}

func InsertEquipamento(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o body da request"))
		return
	}

	var equipamento models.Equipamento
	if erro = json.Unmarshal(requestBody, &equipamento); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter equipamento"))
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao database"))
		return
	}

	defer db.Close()
	equipamentoRepository := repositories.NewRepositoryEquipamento(db)
	equipamento.ID, erro = equipamentoRepository.InsertEquipamento(equipamento)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusCreated, equipamento)

}

func UpdateEquipamento(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro ID"))
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o body da request"))
		return
	}

	var equipamento models.Equipamento
	if erro = json.Unmarshal(requestBody, &equipamento); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter equipamento"))
		return
	}

	db, erro := database.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao database"))
		return
	}

	defer db.Close()
	equipamentoRepository := repositories.NewRepositoryEquipamento(db)
	if erro = equipamentoRepository.UpdateEquipamento(ID, equipamento); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

func DeleteEquipmentoByID(w http.ResponseWriter, r *http.Request) {
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

	equipamentoRepository := repositories.NewRepositoryEquipamento(db)
	if erro = equipamentoRepository.DeleteEquipmentoByID(ID); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}
