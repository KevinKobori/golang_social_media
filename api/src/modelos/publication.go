package modelos

import (
	"errors"
	"strings"
	"time"
)

// Publication representa uma publicação feita por um usuário
type Publication struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
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
	if publication.Title == "" {
		return errors.New("O título é obrigatório e não pode estar em branco")
	}

	if publication.Conteudo == "" {
		return errors.New("O conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publication *Publication) formatar() {
	publication.Title = strings.TrimSpace(publication.Title)
	publication.Conteudo = strings.TrimSpace(publication.Conteudo)
}
