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
    UserRepository *repository.UserRepository
    TokenRepository *repository.TokenRepository
    SessionService *services.SessionService
    TokenService *services.TokenService
    GitHubService *services.GitHubService
}
func (self Controller) Render(w *http.ResponseWriter, data interface{},  pageNames ...string) {
    aux := []string{"views/base.html"}
    for _, fileName := range pageNames {
        aux = append(aux, "views/pages/" + fileName)
    }
    tmplPage, err := template.ParseFiles(aux...)
    if err != nil {
        self.InternalServerError(*w, nil, err)
        return
    }
    tmpl := template.Must(tmplPage, err)

    tmpl.Execute(*w, data)
}
func (self Controller) RenderPage(w http.ResponseWriter, data interface{}, pageName string) {
    pageName = "views/pages/" + pageName
    tmpl, err := template.ParseFiles(pageName)
    if err != nil {
        self.InternalServerError(w, nil, err)
        return
    }
    w.Header().Set("Content-Type", "text/html")
    if err := tmpl.Execute(w, data); err != nil {
        self.InternalServerError(w, nil, err)
    }
}
func (self Controller) Index(w http.ResponseWriter, r *http.Request) {
    sessionExists, err := self.SessionService.SessionExists(r)
    if err != nil {
        fmt.Println("Aqui, erro sessao nao existe")
        self.InternalServerError(w, r, err)
        return
    }
    if sessionExists {
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return
    }
    data := map[string]interface{}{
        "Link": self.GitHubService.GetAuthUri(),
    }
    self.Render(&w, data, "login.html")
    return
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
    userSession, err := self.SessionService.GetSession(r)
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
    user, err := self.UserRepository.GetUser(userSession.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tokens, err := self.TokenRepository.GetTokens(user)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    for _, token := range tokens {
        token.TokenHash = string(token.TokenHash[0:10])
    }
    data := map[string]interface{}{
        "User": user,
        "Tokens": tokens,
        "Csrf": userSession.CsrfToken.Token,
    }
    self.Render(&w, data, "dashboard.html", "tokens.html", "tokensList.html")
    return
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
func (self Controller) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
    fmt.Println(err)
    w.WriteHeader(500)
    return
}
