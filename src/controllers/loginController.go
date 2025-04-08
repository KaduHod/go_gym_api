package controllers

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	"KaduHod/muscles_api/src/services"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type LoginController struct {
   GitHubService *services.GitHubService
   UserService *services.UserService
   Redis *database.Redis
}
func (self LoginController) Auth(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    accessToken, err := self.GitHubService.GetUserToken(code)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    var user core.ApiUser
    user, err = self.GitHubService.GetUserDetails(accessToken)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    exists, err := self.UserService.Exists(user.Login)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    if !exists {
        if err := self.UserService.CreateUser(user); err != nil {
            fmt.Println(err)
            w.WriteHeader(500)
            return
        }
    }
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    id := uuid.New()
    sessionIdCookie := &http.Cookie{
        Name: "session_id",
        Value: id.String(),
        Path: "/",
        HttpOnly: true,
        Secure: false,
        MaxAge: 3600*2,
    }
    http.SetCookie(w, sessionIdCookie)
    self.Redis.Conn.Set(context.Background(), "uuid:"+id.String(), user.Login, time.Hour * 2)
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
func (self LoginController) getLink () string {
    redirectUri := os.Getenv("GITHUB_REDIRECT_URL")
    clientId := os.Getenv("GITHUB_CLIENT_ID")
    loginLink := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=user&redirect_uri=%s", clientId, redirectUri)
    return loginLink
}

func (self LoginController) Index(w http.ResponseWriter, r *http.Request) {
    sessionId, err := r.Cookie("session_id")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    login := self.Redis.Conn.Get(context.Background(), "uuid:" + sessionId.Value).Val()
    exists, err := self.UserService.Exists(login)
    if exists {
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return
    }
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    tmpl, err := template.ParseFiles("src/views/login.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    data := map[string]interface{}{
        "Link": self.getLink(),
    }
    tmpl.Execute(w, data)
}
