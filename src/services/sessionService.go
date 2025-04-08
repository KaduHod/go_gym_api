package services

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type SessionService struct {
    Redis *database.Redis
}
func (self SessionService) NewSession(w *http.ResponseWriter, user core.ApiUser) error {
    id := uuid.New()
    sessionIdCookie := &http.Cookie{
        Name: "session_id",
        Value: id.String(),
        Path: "/",
        HttpOnly: true,
        Secure: false,
        MaxAge: 3600*2,
    }
    http.SetCookie(*w, sessionIdCookie)
    if err := self.Redis.Conn.Set(context.Background(), "uuid:"+id.String(), user.Login, time.Hour * 2).Err(); err != nil {
        fmt.Println("Erro ao criar sessao")
        fmt.Println(err)
        return err
    }
    return nil
}
func (self SessionService) GetUserFromSession(r *http.Request) (string, error) {
    var login string
    sessionId, err := r.Cookie("session_id")
    if err != nil {
        return login, err
    }
    cmd := self.Redis.Conn.Get(context.Background() ,"uuid:"+sessionId.Value)
    if cmd.Err() != nil {
        return login, cmd.Err()
    }
    return cmd.Val(), nil
}
func (self SessionService) SessionExists(r *http.Request) (bool, error) {
    _, err := r.Cookie("session_id")
    if err != nil && err.Error() != "http: named cookie not present" {
        return false, err
    }
    _, err = self.GetUserFromSession(r)
    if err != nil {
        return false, err
    }
    return true, nil
}
