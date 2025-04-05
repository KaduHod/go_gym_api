package main

import (
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
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
    controller := controllers.Controller{
        Db: db,
    }
    musculoSkeletalController := controllers.MusculoSkeletalController{ controller }
    http.HandleFunc("/api/v1/muscles/groups", musculoSkeletalController.ListMuscleGroups)
    http.HandleFunc("/api/v1/muscles/portions", musculoSkeletalController.ListMusclePortions)
    http.HandleFunc("/api/v1/muscles", musculoSkeletalController.ListMuscles)
    http.HandleFunc("/api/v1/joints", musculoSkeletalController.ListJoints)
    http.HandleFunc("/api/v1/moviments", musculoSkeletalController.ListMoviments)
    http.HandleFunc("/api/v1/musculoSkeletalSystem", musculoSkeletalController.ListAmm)
    http.ListenAndServe(":3005", nil)
}
