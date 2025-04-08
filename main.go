package main

import (
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	"KaduHod/muscles_api/src/services"
	"log"
	"net/http"
    _ "KaduHod/muscles_api/docs"
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
    musclesService := services.MuscleService{Db: db}
    movementService := services.MovementService{Db: db}
    jointService := services.JointService{Db: db}
    ammService := services.AmmService{Db: db}
    githubService := services.GitHubService{}
    userService := services.UserService{Db: db}
    sessionService := services.SessionService{Redis: redis}
    controller := controllers.Controller{
        Redis: redis,
        UserService: &userService,
    }
    musculoSkeletalController := controllers.MusculoSkeletalController{
        Controller: controller,
        MuscleService: &musclesService,
        MovementService: &movementService,
        JointService: &jointService,
        AmmService: &ammService,
    }
    loginController := controllers.LoginController{
        GitHubService: &githubService,
        SessionService: &sessionService,
        UserService: &userService,
    }
    server := http.NewServeMux()
    server.HandleFunc("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
    server.HandleFunc("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
    server.HandleFunc("/api/v1/muscles/movement-map", musculoSkeletalController.ListAmm)
    server.HandleFunc("/api/v1/muscles", musculoSkeletalController.ListMuscles)
    server.HandleFunc("/api/v1/joints", musculoSkeletalController.ListJoints)
    server.HandleFunc("/api/v1/movements", musculoSkeletalController.ListMoviments)
    server.HandleFunc("/docs/", httpSwagger.WrapHandler)
    server.HandleFunc("/", loginController.Index)
    server.HandleFunc("/auth/github", loginController.Auth)
    server.HandleFunc("/dashboard", controller.Dashboard)
    http.ListenAndServe(":3005", server)
}
