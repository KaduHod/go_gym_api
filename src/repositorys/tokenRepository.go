package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)
type TokenRepository struct {
    Db *sql.DB
}

func (self TokenRepository) GetTokens(user core.ApiUser) ([]core.UserAPIToken, error) {
    query := `SELECT id, token_name, token_hash, created_at, user_id FROM user_api_tokens WHERE user_id = ? AND deleted_at IS NULL`
    rows, err := self.Db.Query(query, user.Id)
    var resources []core.UserAPIToken
    if err != nil {
        if err == sql.ErrNoRows {
            return resources, nil
        }
        return resources, err
    }
    for rows.Next() {
        var resource core.UserAPIToken
        if err := rows.Scan(&resource.Id, &resource.TokenName, &resource.TokenHash, &resource.CreatedAt, &resource.UserId); err != nil {
            return resources, err
        }
        resources = append(resources, resource)
    }
    return resources, nil
}
func (self TokenRepository) GetTokensByLogin(login string) ([]core.UserAPIToken, error) {
    query := `SELECT tk.id, tk.token_name, tk.token_hash, tk.created_at, tk.user_id FROM user_api_tokens tk
INNER JOIN api_users au ON au.id = tk.user_id
WHERE au.login = ? `
    rows, err := self.Db.Query(query, login)
    var resources []core.UserAPIToken
    if err != nil {
        if err == sql.ErrNoRows {
            return resources, nil
        }
        return resources, err
    }
    for rows.Next() {
        var resource core.UserAPIToken
        if err := rows.Scan(&resource.Id, &resource.TokenName, &resource.TokenHash, &resource.CreatedAt, &resource.UserId); err != nil {
            return resources, err
        }
        resources = append(resources, resource)
    }
    return resources, nil
}
func (self TokenRepository) SaveToken(token core.UserAPIToken) (int64, error) {
    query := `INSERT INTO user_api_tokens (token_name, token_hash, user_id) VALUES (?, ?, ?)`
    result, err := self.Db.Exec(query, token.TokenName, token.TokenHash, token.UserId)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    return id, err
}
func (self TokenRepository) DeleteToken(id int64) error {
    query := `DELETE FROM user_api_tokens WHERE id = ?`
    _, err := self.Db.Exec(query, id)
    return err
}
