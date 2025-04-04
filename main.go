package main

import (
	"KaduHod/muscles_api/src/controllers"
	"KaduHod/muscles_api/src/database"
	"net/http"
)

func main() {
    db := database.ConnetionMysql()
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
