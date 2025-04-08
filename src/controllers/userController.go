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
}

func (self UserController) ListToken(w http.ResponseWriter, r *http.Request) {

}
