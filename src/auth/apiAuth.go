package auth

import (
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"fmt"
	"net/http"
	"strings"
)
type ApiAuthService struct {
    TokenRepository *repository.TokenRepository
    TokenService *services.TokenService
}
func (self *ApiAuthService) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenBearer := r.Header.Get("Authorization")
        if tokenBearer == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        if !strings.Contains(tokenBearer, "Bearer") {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        tokenBearer = strings.Replace(tokenBearer, "Bearer ", "", 1)
        login, err := self.TokenService.GetUserFromToken(tokenBearer)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        tokens, err := self.TokenRepository.GetTokensByLogin(login)
        if err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        var valid bool
        for _, token := range tokens {
            fmt.Println(token)
            valid = self.TokenService.ValidateToken(tokenBearer, token.TokenHash)
            if valid {
                break
            }
        }
        if !valid {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
