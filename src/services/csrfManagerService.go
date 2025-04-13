package services

import (
	"KaduHod/muscles_api/src/core"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

type CsrfService struct {
	cookieName       string
	headerName       string
	cookieMaxAge     int
	secureCookie     bool
	sessionStorage   map[string]core.CsrfTokenInfo
	storageMutex     sync.RWMutex
	cleanupInterval  time.Duration
	tokenLength      int
    sessionService *SessionService
}
func NewCsrfService(sessionService *SessionService) *CsrfService {
	manager := &CsrfService{
		cookieName:      "csrf_token",
		headerName:      "X-CSRF-Token",
		cookieMaxAge:    3600, // 1 hora
		secureCookie:    true,
		sessionStorage:  make(map[string]core.CsrfTokenInfo),
		cleanupInterval: 10 * time.Minute,
		tokenLength:     32,
        sessionService: sessionService,
	}

	return manager
}

// generateToken cria um novo token CSRF seguro
func (m *CsrfService) generateToken() (string, error) {
	bytes := make([]byte, m.tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	hash := sha256.Sum256(bytes)
	token := base64.URLEncoding.EncodeToString(hash[:])
	return token, nil
}

type Middleware func(http.Handler) http.Handler

func (m *CsrfService) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Para requisições GET, HEAD, OPTIONS - apenas gerar token
		if r.Method == "GET" || r.Method == "HEAD" || r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}
        sessionExists, err := m.sessionService.SessionExists(r)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        if !sessionExists {
            http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
        }
		// Verificar token do cabeçalho ou formulário
		requestToken := r.Header.Get(m.headerName)
		if requestToken == "" {
			// Tentar obter do formulário
			requestToken = r.FormValue("csrf_token")
		}

		if !m.validateToken(r, requestToken) {
			http.Error(w, "Invalid CSRF token", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func (m *CsrfService) CreateToken(w http.ResponseWriter) core.CsrfTokenInfo {
    token, err := m.generateToken()
    var tokenInfo core.CsrfTokenInfo
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return tokenInfo
    }
    cookie := &http.Cookie{
        Name:     m.cookieName,
        Value:    token,
        Path:     "/",
        MaxAge:   m.cookieMaxAge,
        HttpOnly: true,
        Secure:   m.secureCookie,
		//Secure:   m.secureCookie,
		//SameSite: http.SameSiteStrictMode,
    }
    http.SetCookie(w, cookie)
    tokenInfo = core.CsrfTokenInfo{
		Token:      token,
		Expiration: time.Now().Add(time.Duration(m.cookieMaxAge) * time.Second),
	}
    return tokenInfo
}
func (m *CsrfService) validateToken(r *http.Request, token string) bool {
    session, err := m.sessionService.GetSession(r)
    if err != nil {
        return false
    }
	// Verificar expiração
	if time.Now().After(session.CsrfToken.Expiration) {
		return false
	}
	// Comparar tokens
	return session.CsrfToken.Token == token
}
