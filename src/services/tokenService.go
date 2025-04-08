package services

import (
	"KaduHod/muscles_api/src/core"
	repository "KaduHod/muscles_api/src/repositorys"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"sync"
)

type TokenService struct {
	TokenRepository *repository.TokenRepository
	UserRepository  *repository.UserRepository
	secret          []byte
	once            sync.Once
}

func NewTokenService(userRepo *repository.UserRepository, tokenRepo *repository.TokenRepository) *TokenService {
	return &TokenService{
		UserRepository:  userRepo,
		TokenRepository: tokenRepo,
	}
}

func (self *TokenService) getSecret() []byte {
	self.once.Do(func() {
		self.secret = []byte(os.Getenv("TOKEN_SECRET_APP"))
	})
	return self.secret
}

// GenerateToken cria um token API no formato: {random}:{userID}
func (self *TokenService) GenerateToken(user core.ApiUser) (string, string, error) {
	// Gerar uma parte aleatória segura usando crypto/rand (mais seguro que math/rand)
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", "", err
	}
	randomPart := base64.URLEncoding.EncodeToString(randomBytes)

	// Codificar o identificador do usuário
	userIdent := base64.URLEncoding.EncodeToString([]byte(user.Login))

	// Criar o token formato: randomPart:userIdent
	token := fmt.Sprintf("%s:%s", randomPart, userIdent)

	// Gerar o hash HMAC para armazenamento (mais rápido que bcrypt e suficiente para tokens)
	hash := self.hashToken(token)

	return token, hash, nil
}

// hashToken gera um hash HMAC do token para armazenamento seguro
func (self *TokenService) hashToken(token string) string {
	h := hmac.New(sha256.New, self.getSecret())
	h.Write([]byte(token))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// GetUserFromToken extrai o usuário do token
func (self *TokenService) GetUserFromToken(token string) (string, error) {
	parts := strings.Split(token, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("formato de token inválido")
	}

	// Decodificar o identificador do usuário
	userIdentBytes, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	return string(userIdentBytes), nil
}

// ValidateToken verifica se um token é válido comparando com o hash armazenado
func (self *TokenService) ValidateToken(token string, storedHash string) bool {
	computedHash := self.hashToken(token)
	return hmac.Equal([]byte(computedHash), []byte(storedHash))
}
