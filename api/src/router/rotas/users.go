package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsers = []Rota{
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
