package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)

type UserRepository struct {
    Db *sql.DB
}

func (s *UserRepository) CreateUser(user core.ApiUser) error {
	query := `INSERT INTO api_users (
		login, avatar_url, url, tipo, nome, empresa, blog, localizacao, email
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := s.Db.Exec(query,
		user.Login,
		user.AvatarURL,
		user.URL,
		user.Tipo,
		user.Nome,
		user.Empresa,
		user.Blog,
		user.Localizacao,
		user.Email,
	)
	return err
}

func (s *UserRepository) GetUser(login string) (*core.ApiUser, error) {
	query := `SELECT
		id, login, avatar_url, url, tipo, nome, empresa, blog, localizacao, email
	FROM api_users WHERE login = ?`
	row := s.Db.QueryRow(query, login)
	user := &core.ApiUser{}
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.AvatarURL,
		&user.URL,
		&user.Tipo,
		&user.Nome,
		&user.Empresa,
		&user.Blog,
		&user.Localizacao,
		&user.Email,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (s *UserRepository) Exists(login string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM api_users WHERE login = ?)`
	var exists bool
	err := s.Db.QueryRow(query, login).Scan(&exists)
	return exists, err
}
