package middlewares

import (
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		proximaFuncao(w, r)
	}
}

// Erro ao inicializar a primeira vez a verificacao do token, pq?
// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if erro := autenticacao.ValidarToken(r); erro != nil {
		// 	respostas.Erro(w, http.StatusUnauthorized, erro)
		// 	return
		// }
		proximaFuncao(w, r)
	}
}
