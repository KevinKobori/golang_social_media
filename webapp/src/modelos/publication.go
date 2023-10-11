package modelos

import "time"

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
