package main

import (
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(".env"); err != nil {
        log.Fatal(err)
    }
    db := database.ConnetionMysql()
    fmt.Println(os.Getenv("DATABASE_HOST"))
    defer db.Close()
    controller := controllers.Controller{
        Db: db,
    }
    muscleController := controllers.MusclesController{ controller }

    http.HandleFunc("/api/v1/muscles/groups", muscleController.ListMuscleGroups)
    http.HandleFunc("/api/v1/muscles/portions", muscleController.ListMusclePortions)
    http.HandleFunc("/api/v1/muscles", muscleController.ListMuscles)
    http.ListenAndServe(":3005", nil)
}
