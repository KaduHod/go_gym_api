package controllers

import (
	"KaduHod/muscles_api/src/services"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type LoginController struct {
   GitHubService services.GitHubService
}
func (self LoginController) Auth(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    accessToken, err := self.GitHubService.GetUserToken(code)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(400)
        return
    }
    w.WriteHeader(200)
    w.Write([]byte(accessToken))
}
func (self LoginController) getLink () string {
    redirectUri := os.Getenv("GITHUB_REDIRECT_URL")
    clientId := os.Getenv("GITHUB_CLIENT_ID")
    loginLink := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=user&redirect_uri=%s", clientId, redirectUri)
    return loginLink
}

func (self LoginController) Index(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("src/views/login.html")
    if err != nil {
        fmt.Println(err)
        w.Write([]byte("<h1>Error :: Contact the admin</h1>"))
        return
    }
    data := map[string]interface{}{
        "Link": self.getLink(),
    }
    tmpl.Execute(w, data)
}
