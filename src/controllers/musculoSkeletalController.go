package controllers

import (
	"KaduHod/muscles_api/src/services"
	"database/sql"
	"fmt"
	"net/http"
)
//carlos
type MusculoSkeletalController struct {
    Controller
    MovementService *services.MovementService
    MuscleService *services.MuscleService
    JointService *services.JointService
    AmmService *services.AmmService

}
// ListMuscleGroups godoc
// @Summary List all muscle groups
// @Description Get a list of all muscle groups without their portions
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.MuscleGroup,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /muscles/groups [get]
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
// ListMusclePortions godoc
// @Summary List all muscle portions
// @Description Get a list of all muscle portions with their group IDs
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.MusclePortion,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /muscles/portions [get]
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
// ListMuscles godoc
// @Summary List all muscles with portions
// @Description Get a hierarchical list of all muscle groups with their portions
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.MuscleGroup,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /muscles [get]
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
    SuccessResponse(w, *resources, len(*resources))
    return
}
// ListJoints godoc
// @Summary List all joints
// @Description Get a list of all joints
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.Joint,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /joints [get]
func (self MusculoSkeletalController) ListJoints(w http.ResponseWriter, r *http.Request) {
    resources, err := self.JointService.GetAll()
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
// ListMoviments godoc
// @Summary List all movements
// @Description Get a list of all possible movements
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.Movement,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /movements [get]
func (self MusculoSkeletalController) ListMoviments(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MovementService.GetAll()
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
// ListAmm godoc
// @Summary List all musculoSkeletalSystem
// @Description Get a list of all muscle-movement-joint relationships with optional filtering
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Param muscle_group query string false "Filter by muscle group name"
// @Param muscle_portion query string false "Filter by muscle portion name"
// @Param joint query string false "Filter by joint/articulation name"
// @Param moviment query string false "Filter by movement name"
// @Success 200 {object} controllers.ResponseSwegger{data=[]core.MuscleMovementInfo,metaData=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseSwegger{data=string}
// @Router /muscles/movement-map [get]
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
