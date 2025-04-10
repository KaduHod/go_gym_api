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
            self.Unauthorized(w)
            return
        }
        if !strings.Contains(tokenBearer, "Bearer") {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        tokenBearer = strings.Replace(tokenBearer, "Bearer ", "", 1)
        login, err := self.TokenService.GetUserFromToken(tokenBearer)
        if err != nil {
            if strings.Contains(err.Error(),"illegal base64") || err.Error() == "formato de token inválido" {
                self.Unauthorized(w)
            } else {
                w.WriteHeader(http.StatusInternalServerError)
            }
            fmt.Println(err)
            return
        }
        tokens, err := self.TokenRepository.GetTokensByLogin(login)
        if err != nil {
            self.Unauthorized(w)
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
            self.Unauthorized(w)
            return
        }
        next.ServeHTTP(w, r)
    })
}
func (self ApiAuthService) Unauthorized(w http.ResponseWriter) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusUnauthorized)
    w.Write([]byte(`{"message":"Unauthorized"}`))
}
