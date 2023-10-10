package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotaSignIn = Rota{
	URI:                "/signIn",
	Metodo:             http.MethodPost,
	Funcao:             controllers.SignIn,
	RequerAutenticacao: false,
}
