package servidor

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar os usuários!"))
		return
	}
	defer rows.Close()

	var usuarios []usuario
	for rows.Next() {
		var usuario usuario

		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)

	// transforma o slice de usuarios em um json para response de um http,
	// para response de http nao se usa o marshal
	if err := json.NewEncoder(w).Encode(usuarios); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os usuarios para json!"))
		return
	}

}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter o parametro para inteiro!"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao conectar com o banco de dados!"))
		return
	}
	defer db.Close()

	row, err := db.Query("select * from usuarios u where u.id = ?", ID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao buscar o usuario!"))
		return
	}

	var usuario usuario
	if row.Next() {
		if err := row.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Erro ao escanear o usuário!"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Usuario não encontrado!"))
		return
	}

	w.WriteHeader(http.StatusOK)

	// transforma o slice de usuarios em um json para response de um http,
	// para response de http nao se usa o marshal
	if err := json.NewEncoder(w).Encode(usuario); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Erro ao converter os usuarios para json!"))
		return
	}
}
