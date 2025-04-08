package controllers

import (
	"KaduHod/muscles_api/src/core"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"html/template"
	"net/http"
)

type LoginController struct {
    Controller
    GitHubService *services.GitHubService
    UserRepository *repository.UserRepository
    SessionService *services.SessionService
}
func (self LoginController) Auth(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    accessToken, err := self.GitHubService.GetUserToken(code)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    var user core.ApiUser
    user, err = self.GitHubService.GetUserDetails(accessToken)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    exists, err := self.UserRepository.Exists(user.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if !exists {
        if err := self.UserRepository.CreateUser(user); err != nil {
            self.InternalServerError(w, r, err)
            return
        }
    }
    if err := self.SessionService.NewSession(&w, user, accessToken); err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    self.Controller.Dashboard(w, r)
}
func (self LoginController) LoggedIndex(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if !sessionExists {
        self.InternalServerError(w, r, err)
        return
    }
    tmpl, err := template.ParseFiles("src/views/logged.html")
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tmpl.Execute(w, nil)
}
func (self LoginController) Index(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if sessionExists {
        self.LoggedIndex(w, r)
        return
    }
    tmpl, err := template.ParseFiles("src/views/login.html")
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    data := map[string]interface{}{
        "Link": self.GitHubService.GetAuthUri(),
    }
    tmpl.Execute(w, data)
}
