package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublications = []Rota{
	{
		URI:                  "/publications",
		Metodo:               http.MethodPost,
		Funcao:               controllers.CriarPublication,
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
	{
		URI:                  "/publications/{publicationId}/atualizar",
		Metodo:               http.MethodGet,
		Funcao:               controllers.CarregarPaginaDeAtualizacaoDePublication,
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
}
