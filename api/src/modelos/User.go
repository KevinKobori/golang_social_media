package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar o usuário recebido
func (user *User) Preparar(etapa string) error {

	if erro := user.validar(etapa); erro != nil {
		return erro
	}

	if erro := user.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validar(etapa string) error {
	if user.Name == "" {
		return errors.New("O name é obrigatório e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && user.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) formatar(etapa string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(user.Senha)
		if erro != nil {
			return erro
		}

		user.Senha = string(senhaComHash)
	}

	return nil
}
