package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
)
type GitHubService struct {}
type gitHubAccessTokenBody struct {
    ClientId string `json:"client_id"`
    ClientSecret string `json:"client_secret"`
    Code string `json:"code"`
    RedirectUri string `json:"redirect_uri"`
}
func (self GitHubService) GetUserToken(code string) (string, error) {
    body := gitHubAccessTokenBody {
        ClientId: os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        Code: code,
        RedirectUri: os.Getenv("GITHUB_REDIRECT_URL"),
    }
    jsonValue, err := json.Marshal(body)
    if err != nil {
        return "", err
    }
    bodyBuffer := bytes.NewBuffer(jsonValue)
    request, err := http.Post("https://github.com/login/oauth/access_token", "application/json", bodyBuffer)
    if err != nil {
        return "", err
    }
    if request.StatusCode != 200 {
        return "", errors.New("Code is not 200")
    }
    bodyBytes, err := io.ReadAll(request.Body)
    if err != nil {
        return "", err
    }
    values, err := url.ParseQuery(string(bodyBytes))
    if err != nil {
        return "", err
    }
    accessToken := values.Get("access_token")
    return accessToken, nil
}
