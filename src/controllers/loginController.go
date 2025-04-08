package controllers

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/services"
	"fmt"
	"html/template"
	"net/http"
)

type LoginController struct {
   GitHubService *services.GitHubService
   UserService *services.UserService
   SessionService *services.SessionService
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
    if err := self.SessionService.NewSession(&w, user); err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
func (self LoginController) LoggedIndex(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if !sessionExists {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    tmpl, err := template.ParseFiles("src/views/logged.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    tmpl.Execute(w, nil)
}
func (self LoginController) Index(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    if sessionExists {
        self.LoggedIndex(w, r)
        return
    }
    tmpl, err := template.ParseFiles("src/views/login.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    data := map[string]interface{}{
        "Link": self.GitHubService.GetAuthUri(),
    }
    tmpl.Execute(w, data)
}
