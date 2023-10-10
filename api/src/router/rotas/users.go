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
		URI:                "/users/{userId}/follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FollowUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/parar-de-follow",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeFollowUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/followers",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarFollowers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/following",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarFollowing,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
