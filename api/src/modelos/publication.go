package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publication representa uma publicação feita por um usuário
type Publication struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

// Preparar vai chamar os métodos para validar e formatar a publicação recebida
func (publication *Publication) Preparar() error {
	if erro := publication.validar(); erro != nil {
		return erro
	}

	publication.formatar()
	return nil
}

func (publication *Publication) validar() error {
	if publication.Titulo == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publication.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publication *Publication) formatar() {
	publication.Titulo = strings.TrimSpace(publication.Titulo)
	publication.Conteudo = strings.TrimSpace(publication.Conteudo)
}
