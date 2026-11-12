// LABORATÓRIO DE SAST: código vulnerável de propósito. Não usar em produção.
package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"
)

// G5: senha fixa no código
const dbPassword = "senha-fixa-no-codigo-123"

var db *sql.DB

// G1: injeção de SQL (entrada do usuário concatenada na consulta)
func buscarUsuario(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	query := "SELECT id, email FROM usuarios WHERE nome = '" + nome + "'"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "erro", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintln(w, "ok")
}

// G-ok: consulta parametrizada, NÃO é falha (serve para medir falso positivo)
func buscarUsuarioSeguro(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	rows, err := db.Query("SELECT id, email FROM usuarios WHERE nome = $1", nome)
	if err != nil {
		http.Error(w, "erro", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	fmt.Fprintln(w, "ok")
}

// G2: injeção de comando
func ping(w http.ResponseWriter, r *http.Request) {
	host := r.URL.Query().Get("host")
	out, _ := exec.Command("sh", "-c", "ping -c 1 "+host).CombinedOutput()
	w.Write(out)
}

// G3: XSS refletido (entrada devolvida sem escapar)
func saudacao(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Olá, %s</h1>", nome)
}

// G4: MD5 para senha
func hashSenha(senha string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(senha)))
}

func main() {
	_ = dbPassword
	_ = hashSenha("teste")
	http.HandleFunc("/usuario", buscarUsuario)
	http.HandleFunc("/usuario-seguro", buscarUsuarioSeguro)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/ola", saudacao)
	http.ListenAndServe(":8080", nil)
}
