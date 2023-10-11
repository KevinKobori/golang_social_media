package modelos

import "time"

// Publication representa uma publicação feita por um usuário
type Publication struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Conteudo   string    `json:"conteudo,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}
