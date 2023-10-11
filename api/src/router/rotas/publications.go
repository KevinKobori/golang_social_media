package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublications = []Rota{
	{
		URI:                  "/publications",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarPublication,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublications,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications/{publicationId}",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublication,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications/{publicationId}",
		Metodo:               http.MethodPut,
		Funcao:               controllers.AtualizarPublication,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications/{publicationId}",
		Metodo:               http.MethodDelete,
		Funcao:               controllers.DeletarPublication,
		RequerAuthentication: true,
	},
	{
		URI:                  "/users/{userId}/publications",
		Metodo:               http.MethodGet,
		Funcao:               controllers.BuscarPublicationsPorUser,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications/{publicationId}/curtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CurtirPublication,
		RequerAuthentication: true,
	},
	{
		URI:                  "/publications/{publicationId}/descurtir",
		Metodo:               http.MethodPost,
		Funcao:               controllers.DescurtirPublication,
		RequerAuthentication: true,
	},
}
