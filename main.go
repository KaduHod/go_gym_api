package main

import (
	_ "KaduHod/muscles_api/docs"
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	repository "KaduHod/muscles_api/src/repositorys"
	"KaduHod/muscles_api/src/services"
	"log"
	"net/http"

	_ "github.com/swaggo/http-swagger" // http-swagger middleware
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

// @title Musculo Eskeletal Api
// @version 1.0
// @description API for Muscles System
// @host gymapi.kadu.tec.br
// @BasePath /api/v1
func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }
    db := database.ConnetionMysql()
    defer db.Close()
    redis := database.NewRedis()
    defer redis.Conn.Close()
    musclesRepository := repository.MuscleRepository{Db: db}
    movementRepository := repository.MovementRepository{Db: db}
    jointRepository := repository.JointRepository{Db: db}
    ammRepository := repository.AmmRepository{Db: db}
    userRepository := repository.UserRepository{Db: db}
    tokenRepository := repository.TokenRepository{Db: db}
    githubService := services.GitHubService{}
    sessionService := services.SessionService{Redis: redis}
    tokenService := services.NewTokenService(&userRepository, &tokenRepository)
    csrfService := services.NewCsrfService(&sessionService)
    controller := controllers.Controller{
        UserRepository: &userRepository,
        SessionService: &sessionService,
        GitHubService: &githubService,
        TokenService: &tokenService,
        TokenRepository: &tokenRepository,
    }
    musculoSkeletalController := controllers.MusculoSkeletalController{
        Controller: controller,
        MuscleRepository: &musclesRepository,
        MovementRepository: &movementRepository,
        JointRepository: &jointRepository,
        AmmRepository: &ammRepository,
    }
    loginController := controllers.LoginController{
        Controller: controller,
        GitHubService: &githubService,
        SessionService: &sessionService,
        UserRepository: &userRepository,
        CsrfService: csrfService,
    }
    server := http.NewServeMux()
    server.HandleFunc("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
    server.HandleFunc("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
    server.HandleFunc("/api/v1/muscles/movement-map", musculoSkeletalController.ListAmm)
    server.HandleFunc("/api/v1/muscles", musculoSkeletalController.ListMuscles)
    server.HandleFunc("/api/v1/joints", musculoSkeletalController.ListJoints)
    server.HandleFunc("/api/v1/movements", musculoSkeletalController.ListMoviments)
    server.HandleFunc("/docs/", httpSwagger.WrapHandler)
    server.HandleFunc("/", controller.Index)
    server.HandleFunc("/auth/github", loginController.Auth)
    server.HandleFunc("/dashboard", controller.Dashboard)
    http.ListenAndServe(":3005", server)
}
