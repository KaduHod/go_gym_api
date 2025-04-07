package controllers

import (
	"KaduHod/muscles_api/src/services"
	"database/sql"
	"fmt"
	"net/http"
)

type MusculoSkeletalController struct {
    Controller
    MovimentService *services.MovimentService
    MuscleService *services.MuscleService
    JointService *services.JointService
    AmmService *services.AmmService

}
func (self MusculoSkeletalController) ListMuscleGroups(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleService.GetAll()
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    SuccessResponse(w, resources, len(resources))
    return
}
func (self MusculoSkeletalController) ListMusclePortions(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleService.GetAllPortions()
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        InternalServerErrorResponse(w, err)
        return
    }
    SuccessResponse(w, resources, len(resources))
    return
}
func (self MusculoSkeletalController) ListMuscles(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleService.GetWithPortions()
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    fmt.Println(resources)
    SuccessResponse(w, *resources, len(*resources))
    return
}
func (self MusculoSkeletalController) ListJoints(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MovimentService.GetAll()
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    SuccessResponse(w, resources, len(resources))
}
func (self MusculoSkeletalController) ListMoviments(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MovimentService.GetAll()
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    SuccessResponse(w, resources, len(resources))
}
func (self MusculoSkeletalController) ListAmm(w http.ResponseWriter, r *http.Request) {
    filters := map[string]string{
        "muscle_group": r.URL.Query().Get("muscle_group"),
        "muscle_portion": r.URL.Query().Get("muscle_portion"),
        "joint": r.URL.Query().Get("joint"),
        "moviment": r.URL.Query().Get("moviment"),
    }
    resources, err := self.AmmService.GetAll(filters)
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    SuccessResponse(w, resources, len(resources))
}
