package equipamento

import (
	"encoding/json"
	"fmt"
	"gofit/banco"
	"gofit/models"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type equipamento struct{}

func NewEquipamento() equipamento {
	var equipamento equipamento
	return equipamento
}

func (e equipamento) GetAll(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM equipamento")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar equipamentos no banco"))
		return
	}

	defer linhas.Close()

	var equipamentos []models.Equipamento

	for linhas.Next() {
		var equipamento models.Equipamento

		if erro := linhas.Scan(&equipamento.ID, &equipamento.Nome, &equipamento.Tipo); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear os equipamentos"))
			return
		}

		equipamentos = append(equipamentos, equipamento)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(equipamentos); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os equipamentos em JSON"))
		return
	}
}

func (e equipamento) GetById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parâmetro ID"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	var equipamento models.Equipamento
	erro = db.QueryRow("SELECT * FROM equipamento WHERE id = $1", ID).Scan(&equipamento.ID, &equipamento.Nome, &equipamento.Tipo)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar equipamentos no banco"))
		fmt.Println(erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(equipamento); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Erro ao converter o equipamento em JSON"))
		return
	}
}

func (e equipamento) Insert(w http.ResponseWriter, r *http.Request) {
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

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO equipamento (nome, tipo) VALUES($1, $2)")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(equipamento.Nome, equipamento.Tipo)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	_, erro = insercao.RowsAffected()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao verificar inserção"))
		fmt.Println(erro)
		return
	}

	message, erro := json.Marshal("Equipamento inserido")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Erro ao converter equipamento em JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(message)
}
func (e equipamento) Update(w http.ResponseWriter, r *http.Request) {
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

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE equipamento SET nome = $2, tipo = $3 WHERE id = $1")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID, equipamento.Nome, equipamento.Tipo); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (e equipamento) Delete(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)

	if handleError(w, erro, "Erro ao converter o parâmetro ID") {
		return
	}

	db, erro := banco.Conectar()

	if handleError(w, erro, "Erro ao conectar ao banco") {
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("DELETE FROM equipamento WHERE id = $1")

	if handleError(w, erro, "Erro ao criar statement") {
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(ID)

	if handleError(w, erro, "Erro ao setar parametros do statement") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func handleError(w http.ResponseWriter, erro error, message string) bool {
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(message))
		fmt.Println(erro)
		return true
	}

	return false

}
