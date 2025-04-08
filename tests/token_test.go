package tests

import (
	"KaduHod/muscles_api/src/core"
	"KaduHod/muscles_api/src/database"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"testing"

	"github.com/joho/godotenv"
)

func TestToken(t *testing.T) {
    godotenv.Load("../.env")
    db := database.ConnetionMysql()
    tokenRepository := repository.TokenRepository{Db: db}
    userRepository := repository.UserRepository{Db: db}
    tokenService := services.NewTokenService(&userRepository, &tokenRepository)
    testUser := core.ApiUser{
        Login: "KaduHod",
    }
    t.Run("Test token creation", func(t *testing.T) {
        _, _, err := tokenService.GenerateToken(testUser)
        if err != nil {
            t.Log(err)
            t.Fail()
        }
    })
    t.Run("Test compare valid token", func(t *testing.T) {
        token, hash, err := tokenService.GenerateToken(testUser)
        if err != nil {
            t.Log(err)
            t.Fail()
        }
        valid := tokenService.ValidateToken(token, hash)
        if !valid {
            t.Log("Invalid token")
            t.Fail()
        }
    })
    t.Run("Test compare invalid token", func(t *testing.T) {
        _, hash, err := tokenService.GenerateToken(testUser)
        if err != nil {
            t.Log(err)
            t.Fail()
        }
        valid := tokenService.ValidateToken("invalid", hash)
        if valid {
            t.Log("Valid token")
            t.Fail()
        }
    })
}
