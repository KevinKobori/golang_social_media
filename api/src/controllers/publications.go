package controllers

import (
	"api/src/authentication"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarPublication adiciona uma nova publicação no banco de dados
func CriarPublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtrairUserID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication modelos.Publication
	if erro = json.Unmarshal(corpoRequisicao, &publication); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	publication.AuthorID = userID

	if erro = publication.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publication.ID, erro = repositorio.Criar(publication)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publication)
}

// BuscarPublications traz as publicações que apareceriam no feed do usuário
func BuscarPublications(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtrairUserID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publications, erro := repositorio.Buscar(userID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publications)
}

// BuscarPublication traz uma única publicação
func BuscarPublication(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publication, erro := repositorio.BuscarPorID(publicationID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publication)
}

// AtualizarPublication altera os dados de uma publicação
func AtualizarPublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtrairUserID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publicationSalvaNoBanco, erro := repositorio.BuscarPorID(publicationID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSalvaNoBanco.AuthorID != userID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar uma publicação que não seja sua"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication modelos.Publication
	if erro = json.Unmarshal(corpoRequisicao, &publication); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publication.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.Atualizar(publicationID, publication); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarPublication exclui os dados de uma publicação
func DeletarPublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtrairUserID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publicationSalvaNoBanco, erro := repositorio.BuscarPorID(publicationID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSalvaNoBanco.AuthorID != userID {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possível deletar uma publicação que não seja sua"))
		return
	}

	if erro = repositorio.Deletar(publicationID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// BuscarPublicationsPorUser traz todas as publicações de um usuário específico
func BuscarPublicationsPorUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	publications, erro := repositorio.BuscarPorUser(userID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, publications)
}

// CurtirPublication adiciona uma curtida na publicação
func CurtirPublication(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	if erro = repositorio.Curtir(publicationID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

// DescurtirPublication subtrai uma curtida na publicação
func DescurtirPublication(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublications(db)
	if erro = repositorio.Descurtir(publicationID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
