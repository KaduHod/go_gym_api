package controllers

import (
	"KaduHod/muscles_api/src/core"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
    Controller
    SessionService *services.SessionService
    UserRepository *repository.UserRepository
    TokenService *services.TokenService
}

func (self UserController) ListTokens(w http.ResponseWriter, r *http.Request) {
    session, err := self.SessionService.GetSession(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    user, err := self.UserRepository.GetUser(session.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tokens, err := self.TokenRepository.GetTokens(user)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tmpl, err := template.ParseFiles("views/pages/tokensList.html")
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    data := map[string]interface{}{
        "Tokens": tokens,
        "Csrf": session.CsrfToken,
    }
    if err := tmpl.ExecuteTemplate(w, "tokensList", data); err != nil {
        self.InternalServerError(w, r, err)
    }
    return
}
func (self UserController) CreateToken(w http.ResponseWriter, r *http.Request) {
    session, err := self.SessionService.GetSession(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    user, err := self.UserRepository.GetUser(session.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tokenName := r.FormValue("token_name")
    if tokenName == "" {
        w.WriteHeader(http.StatusUnprocessableEntity)
        return
    }
    tokens, err := self.TokenRepository.GetTokens(user)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if len(tokens) > 5 {
        self.RenderPage(w, nil, "tokenLimit.html")
        return
    }
    userToken, tokenHash, err := self.TokenService.GenerateToken(user)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    apiToken := core.UserAPIToken{
        TokenName: tokenName,
        TokenHash: tokenHash,
        UserId: user.Id,
    }
    if _, err := self.TokenRepository.SaveToken(apiToken); err != nil {
        w.WriteHeader(http.StatusUnprocessableEntity)
        return
    }
    data := map[string]string{
        "TokenName": tokenName,
        "TokenCreated": userToken,
    }
    tmpl, err := template.ParseFiles("views/pages/tokenCreated.html")
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    tmpl.ExecuteTemplate(w, "tokenCreated",data)
    return
}
func (self UserController) DeleteToken(w http.ResponseWriter, r *http.Request) {
    session, err := self.SessionService.GetSession(r)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    user, err := self.UserRepository.GetUser(session.Login)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    id := chi.URLParam(r, "id")
    tokenIdInt, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if tokenIdInt == 0 {
        w.WriteHeader(http.StatusUnprocessableEntity)
        return
    }
    tokens, err := self.TokenRepository.GetTokens(user)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    for _, token := range tokens {
        if token.Id == tokenIdInt {
            if err := self.TokenRepository.DeleteToken(token.Id); err != nil {
                self.InternalServerError(w, r, err)
                return
            }
            w.WriteHeader(http.StatusOK)
            return
        }
    }
    w.WriteHeader(http.StatusUnprocessableEntity)
    return
}
