package exercicio

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

type exercicio struct{}

func NewExercicio() exercicio {
	var exercicio exercicio
	return exercicio
}

func (e exercicio) GetAll(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	linhas, erro := db.Query("SELECT * FROM exercicio")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar exercícios no banco"))
		return
	}

	defer linhas.Close()

	var exercicios []models.Exercicio

	for linhas.Next() {
		var exercicio models.Exercicio

		if erro := linhas.Scan(&exercicio.ID, &exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID); erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear os exercícios"))
			return
		}

		exercicios = append(exercicios, exercicio)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(exercicios); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os exercícios em JSON"))
		return
	}
}

func (e exercicio) GetById(w http.ResponseWriter, r *http.Request) {
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

	var exercicio models.Exercicio
	erro = db.QueryRow("SELECT * FROM exercicio WHERE id = $1", ID).Scan(&exercicio.ID, &exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar exercicios no banco"))
		fmt.Println(erro)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if erro := json.NewEncoder(w).Encode(exercicio); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Erro ao converter o exercicio em JSON"))
		return
	}
}

func (e exercicio) Insert(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o body da request"))
		return
	}

	var exercicio models.Exercicio
	if erro = json.Unmarshal(requestBody, &exercicio); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter exercicio"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO exercicio (nome, quantidade_minima_repeticoes, quantidade_maxima_repeticoes, id_equipamento) VALUES($1, $2, $3, $4);")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	insercao, erro := statement.Exec(&exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID)

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

	message, erro := json.Marshal("Exercício inserido")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Erro ao converter exercicio em JSON"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(message)
}
func (e exercicio) Update(w http.ResponseWriter, r *http.Request) {
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

	var exercicio models.Exercicio
	if erro = json.Unmarshal(requestBody, &exercicio); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter exercício"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("UPDATE exercicio SET nome= $2, quantidade_minima_repeticoes=$3, quantidade_maxima_repeticoes=$4, id_equipamento=$5 WHERE id = $1")

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID, &exercicio.Nome, &exercicio.QuantidadeMinimaRepeticoes, &exercicio.QuantidadeMaximaRepeticoes, &exercicio.EquipamentoID); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao setar parametros do statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
func (e exercicio) Delete(w http.ResponseWriter, r *http.Request) {
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

	statement, erro := db.Prepare("DELETE FROM exercicio WHERE id = $1")

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
