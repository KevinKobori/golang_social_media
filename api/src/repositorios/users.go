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
		"insert into users (name, nick, email, senha) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

// Buscar traz todos os usuários que atendem um filtro de name ou nick
func (repositorio Users) Buscar(nameOrNickname string) ([]modelos.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname) // %nameOrNickname%

	linhas, erro := repositorio.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNickname, nameOrNickname,
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
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
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
		"select id, name, nick, email, createdAt from users where id = ?",
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
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return modelos.User{}, erro
		}
	}

	return user, nil
}

// Atualizar altera as informações de um usuário no banco de dados
func (repositorio Users) Atualizar(ID uint64, user modelos.User) error {
	statement, erro := repositorio.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
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

// Follow permite que um usuário siga outro
func (repositorio Users) Follow(userID, followerID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil

}

// PararDeFollow permite que um usuário pare de follow o outro
func (repositorio Users) PararDeFollow(userID, followerID uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil

}

// BuscarFollowers traz todos os followers de um usuário
func (repositorio Users) BuscarFollowers(userID uint64) ([]modelos.User, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers s on u.id = s.follower_id where s.user_id = ?`,
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
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil

}

// BuscarFollowing traz todos os usuários que um determinado usuário está following
func (repositorio Users) BuscarFollowing(userID uint64) ([]modelos.User, error) {
	linhas, erro := repositorio.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers s on u.id = s.user_id where s.follower_id = ?`,
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
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
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
