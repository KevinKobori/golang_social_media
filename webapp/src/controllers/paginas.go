package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// CarregarTelaDeSignIn renderiza a tela de signIn
func CarregarTelaDeSignIn(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", 302)
		return
	}

	utils.ExecutarTemplate(w, "signIn.html", nil)
}

// CarregarPaginaDeCadastroDeUser carrega a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUser(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publications", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publications []modelos.Publication
	if erro = json.NewDecoder(response.Body).Decode(&publications); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publications []modelos.Publication
		UserID       uint64
	}{
		Publications: publications,
		UserID:       userID,
	})
}

// CarregarPaginaDeAtualizacaoDePublication carrega a página de edição de publicação
func CarregarPaginaDeAtualizacaoDePublication(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicationID, erro := strconv.ParseUint(parametros["publicationId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publications/%d", config.APIURL, publicationID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publication modelos.Publication
	if erro = json.NewDecoder(response.Body).Decode(&publication); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "atualizar-publication.html", publication)
}

// CarregarPaginaDeUsers carrega a página com os usuários que atendem o filtro passado
func CarregarPaginaDeUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNickname := strings.ToLower(r.URL.Query().Get("user"))
	url := fmt.Sprintf("%s/users?user=%s", config.APIURL, nameOrNickname)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var users []modelos.User
	if erro = json.NewDecoder(response.Body).Decode(&users); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "users.html", users)
}

// CarregarPerfilDoUser carrega a página do perfil do usuário
func CarregarPerfilDoUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	userLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if userID == userLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	user, erro := modelos.BuscarUserCompleto(userID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "user.html", struct {
		User         modelos.User
		UserLogadoID uint64
	}{
		User:         user,
		UserLogadoID: userLogadoID,
	})
}

// CarregarPerfilDoUserLogado carrega a página do perfil do usuário logado
func CarregarPerfilDoUserLogado(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	user, erro := modelos.BuscarUserCompleto(userID, r)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "perfil.html", user)
}

// CarregarPaginaDeEdicaoDeUser carrega a página para edição dos dados do usuário
func CarregarPaginaDeEdicaoDeUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	canal := make(chan modelos.User)
	go modelos.BuscarDadosDoUser(canal, userID, r)
	user := <-canal

	if user.ID == 0 {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: "Erro ao buscar o usuário 2"})
		return
	}

	utils.ExecutarTemplate(w, "edit-user.html", user)
}

// CarregarPaginaDeAtualizacaoDeSenha carrega a página para atualização da senha do usuário
func CarregarPaginaDeAtualizacaoDeSenha(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "atualizar-senha.html", nil)
}
