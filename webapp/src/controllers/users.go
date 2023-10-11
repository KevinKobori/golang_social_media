package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarUser chama a API para cadastrar um usuário no banco de dados
func CriarUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user, erro := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(user))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// PararDeFollowUser chama a API para parar de follow um usuário
func PararDeFollowUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/parar-de-follow", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAuthentication(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// FollowUser chama a API para follow um usuário
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, erro := strconv.ParseUint(parametros["userId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAuthentication(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUser chama a API para editar um usuário
func EditarUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, erro := json.Marshal(map[string]string{
		"name":  r.FormValue("name"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, erro := requisicoes.FazerRequisicaoComAuthentication(r, http.MethodPut, url, bytes.NewBuffer(user))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// AtualizarSenha chama a API para atualizar a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d/atualizar-senha", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAuthentication(r, http.MethodPost, url, bytes.NewBuffer(senhas))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarUser chama a API para deletar um usuário
func DeletarUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	response, erro := requisicoes.FazerRequisicaoComAuthentication(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
