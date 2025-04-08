package controllers

import (
	"KaduHod/muscles_api/src/core"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"fmt"
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
        fmt.Println(err)
        self.Controller.Index(w, r)
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
