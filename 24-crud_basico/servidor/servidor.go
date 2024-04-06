package servidor

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if err = json.Unmarshal(body, &usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o usuário para struct!"))
		return
	}

	// abre conexao com o banco
	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	// prepare statement
	statement, err := db.Prepare("insert into usuarios(nome, email) values(?,?)")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	insert, err := statement.Exec(usuario.Nome, usuario.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	idInserido, err := insert.LastInsertId()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao obter ID inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}
