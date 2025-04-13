package auth

import (
	"KaduHod/muscles_api/src/cache"
	"KaduHod/muscles_api/src/core"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5/middleware"
)
type ApiAuthService struct {
    TokenRepository *repository.TokenRepository
    TokenService *services.TokenService
    CacheService *cache.CacheService
    Log *services.LogService
}
func (self *ApiAuthService) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenBearer := r.Header.Get("Authorization")
        requestID := middleware.GetReqID(r.Context())
        if tokenBearer == "" {
            self.Log.Write("Sem token", requestID)
            self.Unauthorized(w)
            return
        }
        if !strings.Contains(tokenBearer, "Bearer") {
            self.Log.Write("Bearer inválido >> " + tokenBearer + " ||", requestID)
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
            self.Log.Write(err.Error(), requestID)
            return
        }
        var tokens []core.UserAPIToken
        tokens, err = self.CacheService.GetTokensFromUser(login)
        if err != nil {
            self.Log.Write(err.Error(), requestID)
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if len(tokens) == 0 || tokens == nil {
            tokens, err = self.TokenRepository.GetTokensByLogin(login)
            if err != nil {
                self.Log.Write(err.Error(), requestID)
                w.WriteHeader(http.StatusInternalServerError)
                return
            }
            if len(tokens) == 0 {
                self.Unauthorized(w)
                self.Log.Write("Token inválido", requestID)
                return
            }
            go self.CacheService.SetTokensFromUser(login, tokens)
        }
        var valid bool
        for _, token := range tokens {
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
