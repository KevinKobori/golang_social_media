package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasUsers = []Rota{
	{
		URI:                "/create-user",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/search-users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/perfil",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPerfilDoUserLogado,
		RequerAutenticacao: true,
	},
	{
		URI:                "/edit-user",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeEdicaoDeUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/edit-user",
		Metodo:             http.MethodPut,
		Funcao:             controllers.EditarUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizar-senha",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeAtualizacaoDeSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
	{
		URI:                "/deletar-user",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUser,
		RequerAutenticacao: true,
	},
}
