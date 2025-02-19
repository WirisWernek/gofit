package planoTreino

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

type planoTreino struct{}

func NewPlanoTreino() planoTreino {
	var planoTreino planoTreino
	return planoTreino
}

func (p planoTreino) GetAll(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if handleError(w, erro, "Erro ao conectar ao banco") {
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM plano_treino")

	if handleError(w, erro, "Erro ao buscar planos de treino no banco") {
		return
	}

	defer linhas.Close()

	var planoTreinos []models.PlanoTreino

	for linhas.Next() {
		var planoTreino models.PlanoTreino

		erro = linhas.Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo)

		if handleError(w, erro, "Erro ao escanear os planos de treino") {
			return
		}

		planoTreinos = append(planoTreinos, planoTreino)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	erro = json.NewEncoder(w).Encode(planoTreinos)

	if handleError(w, erro, "Erro ao converter os planos de treino em JSON") {
		return
	}
}

func (p planoTreino) GetById(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	// a função ParseUint recebe 3 parâmetros: 1º a variavel a ser convertida, 2º a base utilizada, 3º o tamanho em bits
	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if handleError(w, erro, "Erro ao converter o parâmetro ID") {
		return
	}

	db, erro := banco.Conectar()

	if handleError(w, erro, "Erro ao conectar ao banco") {
		return
	}

	defer db.Close()

	var planoTreino models.PlanoTreino
	erro = db.QueryRow("SELECT * FROM plano_treino WHERE id = $1", ID).Scan(&planoTreino.ID, &planoTreino.Nome, &planoTreino.Descricao, &planoTreino.Inicio, &planoTreino.Descanso, &planoTreino.Ativo)

	if handleError(w, erro, "Erro ao buscar o plano de treino no banco") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	erro = json.NewEncoder(w).Encode(planoTreino)

	if handleError(w, erro, "Erro ao converter o plano de treino em JSON") {
		return
	}
}

func (p planoTreino) Insert(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if handleError(w, erro, "Falha ao ler o body da request") {
		return
	}

	var planoTreino models.PlanoTreino
	erro = json.Unmarshal(requestBody, &planoTreino)

	if handleError(w, erro, "Erro ao converter o plano de treino") {
		return
	}

	db, erro := banco.Conectar()

	if handleError(w, erro, "Erro ao conectar ao banco") {
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO plano_treino (nome, descricao, inicio, descanso, ativo) VALUES($1, $2, $3, $4, $5)")

	if handleError(w, erro, "Erro ao criar statement") {
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo)

	if handleError(w, erro, "Erro ao setar parametros do statement") {
		return
	}

	_, erro = insercao.RowsAffected()

	if handleError(w, erro, "Erro ao verificar inserção") {
		return
	}

	message, erro := json.Marshal("Equipamento inserido")

	if handleError(w, erro, "Erro ao gerar mensagem") {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(message)
}

func (p planoTreino) Update(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 64)

	if handleError(w, erro, "Erro ao converter o parâmetro ID") {
		return
	}

	requestBody, erro := io.ReadAll(r.Body)

	if handleError(w, erro, "Falha ao ler o body da request") {
		return
	}

	var planoTreino models.PlanoTreino
	erro = json.Unmarshal(requestBody, &planoTreino)

	if handleError(w, erro, "Erro ao converter plano de treino") {
		return
	}

	db, erro := banco.Conectar()

	if handleError(w, erro, "Erro ao conectar ao banco") {
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE plano_treino SET nome=$2, descricao=$3, inicio=$4, descanso=$5, ativo=$6 WHERE id=$1")

	if handleError(w, erro, "Erro ao criar statement") {
		return
	}

	defer statement.Close()

	_, erro = statement.Exec(ID, planoTreino.Nome, planoTreino.Descricao, planoTreino.Inicio, planoTreino.Descanso, planoTreino.Ativo)

	if handleError(w, erro, "Erro ao setar parametros do statement") {
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (p planoTreino) Delete(w http.ResponseWriter, r *http.Request) {
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

	statement, erro := db.Prepare("DELETE FROM plano_treino WHERE id = $1")

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
