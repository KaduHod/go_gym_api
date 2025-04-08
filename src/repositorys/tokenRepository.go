package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)
type TokenRepository struct {
    Db *sql.DB
}

func (self TokenRepository) GetTokens(user core.ApiUser) ([]core.UserAPIToken, error) {
    query := `SELECT id, token_name, token_hash, created_at, expires_at, user_id FROM user_api_tokens WHERE user_id = ?`
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
        if err := rows.Scan(&resource.Id, &resource.TokenName, &resource.TokenHash, &resource.CreatedAt, &resource.DeletedAt, &resource.UserId); err != nil {
            return resources, err
        }
        resources = append(resources, resource)
    }
    return resources, nil
}
