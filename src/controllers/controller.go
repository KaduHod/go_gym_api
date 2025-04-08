package controllers

import (
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)
type Controller struct {
    SessionService *services.SessionService
    UserRepository *repository.UserRepository
    GitHubService *services.GitHubService
}
func (self Controller) Render(w *http.ResponseWriter, pageName string, data interface{}) {
    tmplPage, err := template.ParseFiles("views/base.html", "views/pages/" + pageName)
    if err != nil {
        self.InternalServerError(*w, nil, err)
        return
    }
    tmpl := template.Must(tmplPage, err)
    tmpl.Execute(*w, data)
}
func (self Controller) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Println(err)
    w.WriteHeader(500)
    return
}
func (self Controller) Index(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if sessionExists {
        self.Dashboard(w, r)
        return
    }
    data := map[string]interface{}{
        "Link": self.GitHubService.GetAuthUri(),
    }
    self.Render(&w, "login.html", data)
}
func (self Controller) Dashboard(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if !sessionExists {
        self.Index(w, r)
        return
    }
    userSession, err := self.SessionService.GetUserFromSession(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    exists, err := self.UserRepository.Exists(userSession.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if !exists {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    self.Render(&w, "dashboard.html", nil)
}
type MetaData struct {
	TotalItens int `json:"total_itens"`
}

type Response[T any] struct {
	Status   string  `json:"status"`
	Data     T       `json:"data"`
	MetaData MetaData `json:"metadata"`
}
type ResponseSwegger struct {
	Status   string  `json:"status"`
    Data     interface {}       `json:"data"`
	MetaData MetaData `json:"metadata"`
}
func SuccessResponse[T any](w http.ResponseWriter, data T, totalItems int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response[T]{
		Status: "success",
		Data:   data,
		MetaData: MetaData{
			TotalItens: totalItems,
		},
	})
}
func InternalServerErrorResponse(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Response[string]{
		Status: "fail",
		Data:   err.Error(),
	})
}
func EmptyResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(Response[interface{}]{
		Status: "success",
		Data:   nil,
	})
}
