package services

import (
	"KaduHod/muscles_api/src/core"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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
func (self GitHubService) GetUserDetails(accessToken string) (core.ApiUser, error) {
    request, err := http.NewRequest("GET", "https://api.github.com/user", nil)
    var user core.ApiUser
    if err != nil {
        return user, err
    }
    request.Header.Set("Authorization", "token "+accessToken)
    client := &http.Client{}
    response, err := client.Do(request)
    if err != nil {
        return user, err
    }
    if response.StatusCode != 200 {
        return user, errors.New("Code is not 200")
    }
    bodyBytes, err := io.ReadAll(response.Body)
    if err != nil {
        return user, err
    }
    if err := json.Unmarshal(bodyBytes, &user); err != nil {
        return user, err
    }
    return user, nil
}
func (self GitHubService) GetAuthUri () string {
    redirectUri := os.Getenv("GITHUB_REDIRECT_URL")
    clientId := os.Getenv("GITHUB_CLIENT_ID")
    loginLink := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=user&redirect_uri=%s", clientId, redirectUri)
    return loginLink
}
