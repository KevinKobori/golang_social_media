package modelos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
)

// User representa uma pessoa utilizando a rede social
type User struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []User       `json:"seguidores"`
	Seguindo    []User       `json:"seguindo"`
	Publicacoes []Publicacao `json:"publicacoes"`
}

// BuscarUserCompleto faz 4 requisições na API para montar o usuário
func BuscarUserCompleto(userID uint64, r *http.Request) (User, error) {
	canalUser := make(chan User)
	canalSeguidores := make(chan []User)
	canalSeguindo := make(chan []User)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUser(canalUser, userID, r)
	go BuscarSeguidores(canalSeguidores, userID, r)
	go BuscarSeguindo(canalSeguindo, userID, r)
	go BuscarPublicacoes(canalPublicacoes, userID, r)

	var (
		user        User
		seguidores  []User
		seguindo    []User
		publicacoes []Publicacao
	)

	for i := 0; i < 4; i++ {
		select {
		case userCarregado := <-canalUser:
			if userCarregado.ID == 0 {
				return User{}, errors.New("Erro ao buscar o usuário 3")
			}

			user = userCarregado

		case seguidoresCarregados := <-canalSeguidores:
			if seguidoresCarregados == nil {
				return User{}, errors.New("Erro ao buscar os seguidores")
			}

			seguidores = seguidoresCarregados

		case seguindoCarregados := <-canalSeguindo:
			if seguindoCarregados == nil {
				return User{}, errors.New("Erro ao buscar quem o usuário está seguindo")
			}

			seguindo = seguindoCarregados

		case publicacoesCarregadas := <-canalPublicacoes:
			if publicacoesCarregadas == nil {
				return User{}, errors.New("Erro ao buscar as publicações")
			}

			publicacoes = publicacoesCarregadas
		}
	}

	user.Seguidores = seguidores
	user.Seguindo = seguindo
	user.Publicacoes = publicacoes

	return user, nil
}

// BuscarDadosDoUser chama a API para buscar os dados base do usuário
func BuscarDadosDoUser(canal chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if erro = json.NewDecoder(response.Body).Decode(&user); erro != nil {
		canal <- User{}
		return
	}

	canal <- user
}

// BuscarSeguidores chama a API para buscar os seguidores do usuário
func BuscarSeguidores(canal chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/seguidores", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguidores []User
	if erro = json.NewDecoder(response.Body).Decode(&seguidores); erro != nil {
		canal <- nil
		return
	}

	if seguidores == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- seguidores
}

// BuscarSeguindo chama a API para buscar os usuários seguidos por um usuário
func BuscarSeguindo(canal chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/seguindo", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var seguindo []User
	if erro = json.NewDecoder(response.Body).Decode(&seguindo); erro != nil {
		canal <- nil
		return
	}

	if seguindo == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- seguindo
}

// BuscarPublicacoes chama a API para buscar as publicações de um usuário
func BuscarPublicacoes(canal chan<- []Publicacao, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publicacoes", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publicacoes []Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		canal <- nil
		return
	}

	if publicacoes == nil {
		canal <- make([]Publicacao, 0)
		return
	}

	canal <- publicacoes
}
