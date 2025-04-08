package controllers

import (
	"KaduHod/muscles_api/src/database"
	"KaduHod/muscles_api/src/services"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)
type Controller struct {
    Redis *database.Redis
    UserService *services.UserService
}
func (self Controller) Dashboard(w http.ResponseWriter, r *http.Request) {
    sessionId, err := r.Cookie("session_id")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    login := self.Redis.Conn.Get(context.Background(), "uuid:" + sessionId.Value).Val()
    exists, err := self.UserService.Exists(login)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    if !exists {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    user, err := self.UserService.GetUser(login)
    if err != nil {
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }
    tmpl, err := template.ParseFiles("src/views/logged.html")
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(500)
        return
    }
    fmt.Println(user)
    tmpl.Execute(w, nil)
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
