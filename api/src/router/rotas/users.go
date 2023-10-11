package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsers = []Rota{
	{
		URI:                  "/users",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUser,
		RequerAuthentication: false,
	},
	{
		URI:                  "/users",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUsers,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/follow",
		Metodo:               http.MethodPost,
		Funcao:               controllers.FollowUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/parar-de-follow",
		Metodo:               http.MethodPost,
		Funcao:               controllers.PararDeFollowUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/followers",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarFollowers,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/following",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarFollowing,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/atualizar-senha",
		Metodo:               http.MethodPost,
		Funcao:               controllers.AtualizarSenha,
		RequerAuthentication: true,
	},
}
