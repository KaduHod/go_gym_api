package main

import (
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	"KaduHod/muscles_api/src/services"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }
    db := database.ConnetionMysql()
    defer db.Close()
    musclesService := services.MuscleService{Db: db}
    movimentService := services.MovimentService{Db: db}
    jointService := services.JointService{Db: db}
    controller := controllers.Controller{}
    musculoSkeletalController := controllers.MusculoSkeletalController{
        Controller: controller,
        MuscleService: &musclesService,
        MovimentService: &movimentService,
        JointService: &jointService,
    }
    server := http.NewServeMux()
    server.HandleFunc("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
    server.HandleFunc("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
    server.HandleFunc("/api/v1/muscles", musculoSkeletalController.ListMuscles)
    server.HandleFunc("/api/v1/joints", musculoSkeletalController.ListJoints)
    server.HandleFunc("/api/v1/moviments", musculoSkeletalController.ListMoviments)
    server.HandleFunc("/api/v1/musculoSkeletalSystem", musculoSkeletalController.ListAmm)
    http.ListenAndServe(":3005", server)
}
