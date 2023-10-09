package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

// Users representa um repositório de users
type Users struct {
	db *sql.DB
}

// NovoRepositorioDeUsers cria um repositório de usuários
func NovoRepositorioDeUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Users) Criar(user modelos.User) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into users (nome, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(user.Nome, user.Nick, user.Email, user.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Users) Buscar(nomeOuNick string) ([]modelos.User, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from users where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []modelos.User

	for linhas.Next() {
		var user modelos.User

		if erro = linhas.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Users) BuscarPorID(ID uint64) (modelos.User, error) {
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from users where id = ?",
		ID,
	)
	if erro != nil {
		return modelos.User{}, erro
	}
	defer linhas.Close()

	var user modelos.User

	if linhas.Next() {
		if erro = linhas.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return modelos.User{}, erro
		}
	}

	return user, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Users) Atualizar(ID uint64, user modelos.User) error {
	statement, erro := repositorio.db.Prepare(
		"update users set nome = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Nome, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Users) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

// BuscarPorEmail busca um usuário por email e retorna o seu id e senha com hash
func (repositorio Users) BuscarPorEmail(email string) (modelos.User, error) {
	linha, erro := repositorio.db.Query("select id, senha from users where email = ?", email)
	if erro != nil {
		return modelos.User{}, erro
	}
	defer linha.Close()

	var user modelos.User

	if linha.Next() {
		if erro = linha.Scan(&user.ID, &user.Senha); erro != nil {
			return modelos.User{}, erro
		}
	}

	return user, nil

}

// Seguir permite que um usuário siga outro
func (repositorio Users) Seguir(userID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into seguidores (user_id, seguidor_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

// PararDeSeguir permite que um usuário pare de seguir o outro
func (repositorio Users) PararDeSeguir(userID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from seguidores where user_id = ? and seguidor_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, seguidorID); erro != nil {
		return erro
	}

	return nil

}

// BuscarSeguidores traz todos os seguidores de um usuário
func (repositorio Users) BuscarSeguidores(userID uint64) ([]modelos.User, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from users u inner join seguidores s on u.id = s.seguidor_id where s.user_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []modelos.User
	for linhas.Next() {
		var user modelos.User

		if erro = linhas.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil

}

// BuscarSeguindo traz todos os usuários que um determinado usuário está seguindo
func (repositorio Users) BuscarSeguindo(userID uint64) ([]modelos.User, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.nome, u.nick, u.email, u.criadoEm
		from users u inner join seguidores s on u.id = s.user_id where s.seguidor_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var users []modelos.User

	for linhas.Next() {
		var user modelos.User

		if erro = linhas.Scan(
			&user.ID,
			&user.Nome,
			&user.Nick,
			&user.Email,
			&user.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

// BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Users) BuscarSenha(userID uint64) (string, error) {
	linha, erro := repositorio.db.Query("select senha from users where id = ?", userID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var user modelos.User

	if linha.Next() {
		if erro = linha.Scan(&user.Senha); erro != nil {
			return "", erro
		}
	}

	return user.Senha, nil
}

// AtualizarSenha altera a senha de um usuário
func (repositorio Users) AtualizarSenha(userID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("update users set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(senha, userID); erro != nil {
		return erro
	}

	return nil
}
