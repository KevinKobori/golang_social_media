package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publicacoes representa um repositório de publicações
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

// BuscarPorID traz uma única publicação do banco de dados
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linha, erro := repositorio.db.Query(`
	select p.*, u.nick from 
	publicacoes p inner join users u
	on u.id = p.autor_id where p.id = ?`,
		publicacaoID,
	)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacao modelos.Publicacao

	if linha.Next() {
		if erro = linha.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}

// Buscar traz as publicações dos usuários followeds e também do próprio usuário que fez a requisição
func (repositorio Publicacoes) Buscar(userID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p 
	inner join users u on u.id = p.autor_id 
	inner join followers s on p.autor_id = s.user_id 
	where u.id = ? or s.follower_id = ?
	order by 1 desc`,
		userID, userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publicacoes where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUser traz todas as publicações de um usuário específico
func (repositorio Publicacoes) BuscarPorUser(userID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publicacoes p
		join users u on u.id = p.autor_id
		where p.autor_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao

	for linhas.Next() {
		var publicacao modelos.Publicacao

		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return publicacoes, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set curtidas = curtidas + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		update publicacoes set curtidas = 
		CASE 
			WHEN curtidas > 0 THEN curtidas - 1
			ELSE 0 
		END
		where id = ?
	`)
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(publicacaoID); erro != nil {
		return erro
	}

	return nil
}
