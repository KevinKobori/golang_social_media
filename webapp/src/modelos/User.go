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
	ID           uint64        `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Nick         string        `json:"nick"`
	CreatedAt    time.Time     `json:"createdAt"`
	Followers    []User        `json:"followers"`
	Following    []User        `json:"following"`
	Publications []Publication `json:"publications"`
}

// BuscarUserCompleto faz 4 requisições na API para montar o usuário
func BuscarUserCompleto(userID uint64, r *http.Request) (User, error) {
	canalUser := make(chan User)
	canalFollowers := make(chan []User)
	canalFollowing := make(chan []User)
	canalPublications := make(chan []Publication)

	go BuscarDadosDoUser(canalUser, userID, r)
	go BuscarFollowers(canalFollowers, userID, r)
	go BuscarFollowing(canalFollowing, userID, r)
	go BuscarPublications(canalPublications, userID, r)

	var (
		user         User
		followers    []User
		following    []User
		publications []Publication
	)

	for i := 0; i < 4; i++ {
		select {
		case userCarregado := <-canalUser:
			if userCarregado.ID == 0 {
				return User{}, errors.New("Erro ao buscar o usuário 3")
			}

			user = userCarregado

		case followersCarregados := <-canalFollowers:
			if followersCarregados == nil {
				return User{}, errors.New("Erro ao buscar os followers")
			}

			followers = followersCarregados

		case followingCarregados := <-canalFollowing:
			if followingCarregados == nil {
				return User{}, errors.New("Erro ao buscar quem o usuário está following")
			}

			following = followingCarregados

		case publicationsCarregadas := <-canalPublications:
			if publicationsCarregadas == nil {
				return User{}, errors.New("Erro ao buscar as publicações")
			}

			publications = publicationsCarregadas
		}
	}

	user.Followers = followers
	user.Following = following
	user.Publications = publications

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

// BuscarFollowers chama a API para buscar os followers do usuário
func BuscarFollowers(canal chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if erro = json.NewDecoder(response.Body).Decode(&followers); erro != nil {
		canal <- nil
		return
	}

	if followers == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- followers
}

// BuscarFollowing chama a API para buscar os usuários followeds por um usuário
func BuscarFollowing(canal chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/following", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var following []User
	if erro = json.NewDecoder(response.Body).Decode(&following); erro != nil {
		canal <- nil
		return
	}

	if following == nil {
		canal <- make([]User, 0)
		return
	}

	canal <- following
}

// BuscarPublications chama a API para buscar as publicações de um usuário
func BuscarPublications(canal chan<- []Publication, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/publications", config.APIURL, userID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		canal <- nil
		return
	}
	defer response.Body.Close()

	var publications []Publication
	if erro = json.NewDecoder(response.Body).Decode(&publications); erro != nil {
		canal <- nil
		return
	}

	if publications == nil {
		canal <- make([]Publication, 0)
		return
	}

	canal <- publications
}
