package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublications = []Rota{
	{
		URI:                "/publications",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublications,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{publicationId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{userId}/publications",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarPublicationsPorUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{publicationId}/curtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CurtirPublication,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publications/{publicationId}/descurtir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.DescurtirPublication,
		RequerAutenticacao: true,
	},
}
