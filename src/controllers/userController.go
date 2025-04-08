package controllers

import (
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"net/http"
)

type UserController struct {
    Controller
    SessionService *services.SessionService
    UserRepository *repository.UserRepository
    TokenService *services.TokenService
}

func (self UserController) ListToken(w http.ResponseWriter, r *http.Request) {

}
func (self UserController) CreateToken(w http.ResponseWriter, r *http.Request) {

}
