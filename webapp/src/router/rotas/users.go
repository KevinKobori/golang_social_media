package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsers = []Rota{
	{
		URI:                  "/create-user",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeCadastroDeUser,
		RequerAuthentication: false,
	},
	{
		URI:                  "/users",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarUser,
		RequerAuthentication: false,
	},
	{
		URI:                  "/search-users",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeUsers,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPerfilDoUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/parar-de-follow",
		Metodo:               http.MethodPost,
		Funcao:               controllers.PararDeFollowUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/follow",
		Metodo:               http.MethodPost,
		Funcao:               controllers.FollowUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/perfil",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPerfilDoUserLogado,
		RequerAuthentication: true,
	},
	{
		URI:                  "/edit-user",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeEdicaoDeUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/edit-user",
		Metodo:               http.MethodPut,
		Funcao:               controllers.EditarUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/atualizar-senha",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeAtualizacaoDeSenha,
		RequerAuthentication: true,
	},
	{
		URI:                  "/atualizar-senha",
		Metodo:               http.MethodPost,
		Funcao:               controllers.AtualizarSenha,
		RequerAuthentication: true,
	},
	{
		URI:                  "/deletar-user",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarUser,
		RequerAuthentication: true,
	},
}
