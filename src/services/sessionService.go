package services

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/google/uuid"
)
type SessionService struct {
    Redis *database.Redis
}
func (self SessionService) NewSession(w *http.ResponseWriter, user core.ApiUser, githubAccessToken string, csrfToken core.CsrfTokenInfo) error {
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
    value, err := json.Marshal(core.UserSessionData{Login: user.Login, AccessToken: githubAccessToken, CsrfToken: csrfToken})
    if err != nil {
       return err
    }
    if err := self.Redis.Conn.Set(context.Background(), "uuid:"+id.String(), value, time.Hour * 2).Err(); err != nil {
        fmt.Println("Erro ao criar sessao")
        fmt.Println(err)
        return err
    }
    return nil
}
func (self SessionService) GetSession(r *http.Request) (core.UserSessionData, error) {
    var userSessionData core.UserSessionData
    sessionId, err := r.Cookie("session_id")
    if err != nil {
        return userSessionData, err
    }
    cmd := self.Redis.Conn.Get(context.Background() ,"uuid:"+sessionId.Value)
    if cmd.Err() != nil {
        return userSessionData, cmd.Err()
    }
    bytes, err := cmd.Bytes()
    if err != nil {
        return userSessionData, err
    }
    if err := json.Unmarshal(bytes, &userSessionData); err != nil {
        return userSessionData, err
    }
    return userSessionData, nil
}
func (self SessionService) SessionExists(r *http.Request) (bool, error) {
    _, err := r.Cookie("session_id")
    errs := []string{"http: named cookie not present"}
    if err != nil && slices.Contains(errs, err.Error()) {
        return false, nil
    }
    if err != nil && !slices.Contains(errs, err.Error()) {
        return false, err
    }
    return true, nil
}
func (self SessionService) GetSessionId(r *http.Request) (string, error) {
    sessionId, err := r.Cookie("session_id")
    if err != nil {
        return "", err
    }
    return sessionId.Value, nil
}
