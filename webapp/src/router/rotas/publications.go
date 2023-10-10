package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var rotasPublications = []Rota{
	{
		URI:                "/publications",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarPublication,
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
	{
		URI:                "/publications/{publicationId}/atualizar",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeAtualizacaoDePublication,
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
}
