package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Publications representa um repositório de publicações
type Publications struct {
	db *sql.DB
}

// NovoRepositorioDePublications cria um repositório de publicações
func NovoRepositorioDePublications(db *sql.DB) *Publications {
	return &Publications{db}
}

// Criar insere uma publicação no banco de dados
func (repositorio Publications) Criar(publication modelos.Publication) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into publications (title, conteudo, author_id) values (?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publication.Title, publication.Conteudo, publication.AuthorID)
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
func (repositorio Publications) BuscarPorID(publicationID uint64) (modelos.Publication, error) {
	linha, erro := repositorio.db.Query(`
	select p.*, u.nick from 
	publications p inner join users u
	on u.id = p.author_id where p.id = ?`,
		publicationID,
	)
	if erro != nil {
		return modelos.Publication{}, erro
	}
	defer linha.Close()

	var publication modelos.Publication

	if linha.Next() {
		if erro = linha.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Conteudo,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return modelos.Publication{}, erro
		}
	}

	return publication, nil
}

// Buscar traz as publicações dos usuários followeds e também do próprio usuário que fez a requisição
func (repositorio Publications) Buscar(userID uint64) ([]modelos.Publication, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publications p 
	inner join users u on u.id = p.author_id 
	inner join followers s on p.author_id = s.user_id 
	where u.id = ? or s.follower_id = ?
	order by 1 desc`,
		userID, userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publications []modelos.Publication

	for linhas.Next() {
		var publication modelos.Publication

		if erro = linhas.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Conteudo,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Atualizar altera os dados de uma publicação no banco de dados
func (repositorio Publications) Atualizar(publicationID uint64, publication modelos.Publication) error {
	statement, erro := repositorio.db.Prepare("update publications set title = ?, conteudo = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publication.Title, publication.Conteudo, publicationID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui uma publicação do banco de dados
func (repositorio Publications) Deletar(publicationID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from publications where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorUser traz todas as publicações de um usuário específico
func (repositorio Publications) BuscarPorUser(userID uint64) ([]modelos.Publication, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from publications p
		join users u on u.id = p.author_id
		where p.author_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publications []modelos.Publication

	for linhas.Next() {
		var publication modelos.Publication

		if erro = linhas.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Conteudo,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); erro != nil {
			return nil, erro
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

// Curtir adiciona uma curtida na publicação
func (repositorio Publications) Curtir(publicationID uint64) error {
	statement, erro := repositorio.db.Prepare("update publications set likes = likes + 1 where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}

// Descurtir subtrai uma curtida na publicação
func (repositorio Publications) Descurtir(publicationID uint64) error {
	statement, erro := repositorio.db.Prepare(`
		update publications set likes = 
		CASE 
			WHEN likes > 0 THEN likes - 1
			ELSE 0 
		END
		where id = ?
	`)
	if erro != nil {
		return erro
	}

	if _, erro = statement.Exec(publicationID); erro != nil {
		return erro
	}

	return nil
}
