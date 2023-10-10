package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasSignIn = []Rota{
	{
		URI:                "/",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeSignIn,
		RequerAutenticacao: false,
	},
	{
		URI:                "/signIn",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarTelaDeSignIn,
		RequerAutenticacao: false,
	},
	{
		URI:                "/signIn",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FazerSignIn,
		RequerAutenticacao: false,
	},
}
