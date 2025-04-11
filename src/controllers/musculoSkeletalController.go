package controllers

import (
	repository "KaduHod/muscles_api/src/repositorys"
	"database/sql"
	"fmt"
	"net/http"
)

//carlos
type MusculoSkeletalController struct {
    Controller
    MovementRepository *repository.MovementRepository
    MuscleRepository *repository.MuscleRepository
    JointRepository *repository.JointRepository
    AmmRepository *repository.AmmRepository
}
// ListMuscleGroups godoc
// @Summary List all muscle groups
// @Description Get a list of all muscle groups without their portions
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.MuscleGroup,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /muscles/groups [get]
func (self MusculoSkeletalController) ListMuscleGroups(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleRepository.GetAll()
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
    return
}
// ListMusclePortions godoc
// @Summary List all muscle portions
// @Description Get a list of all muscle portions with their group IDs
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.MusclePortion,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /muscles/portions [get]
func (self MusculoSkeletalController) ListMusclePortions(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleRepository.GetAllPortions()
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
    return
}
// ListMuscles godoc
// @Summary List all muscles with portions
// @Description Get a hierarchical list of all muscle groups with their portions
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.MuscleGroup,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /muscles [get]
func (self MusculoSkeletalController) ListMuscles(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MuscleRepository.GetWithPortions()
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(*resources))
    return
}
// ListJoints godoc
// @Summary List all joints
// @Description Get a list of all joints
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.Joint,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /joints [get]
func (self MusculoSkeletalController) ListJoints(w http.ResponseWriter, r *http.Request) {
    resources, err := self.JointRepository.GetAll()
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
}
// ListMoviments godoc
// @Summary List all movements
// @Description Get a list of all possible movements
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.Movement,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /movements [get]
func (self MusculoSkeletalController) ListMoviments(w http.ResponseWriter, r *http.Request) {
    resources, err := self.MovementRepository.GetAll()
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
}
// ListAmm godoc
// @Summary Map muscles, joints e movements.
// @Description Get a list of all muscle-movement-joint relationships with optional filtering
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Param muscle_group query string false "Filter by muscle group name"
// @Param muscle_portion query string false "Filter by muscle portion name"
// @Param joint query string false "Filter by joint name"
// @Param movement query string false "Filter by movement name"
// @Success 200 {object} controllers.ResponseDescription{data=[]core.MuscleMovementInfo,metadata=controllers.MetaData}
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /muscles/movement-map [get]
func (self MusculoSkeletalController) ListAmm(w http.ResponseWriter, r *http.Request) {
    filters := map[string]string{
        "muscle_group": r.URL.Query().Get("muscle_group"),
        "muscle_portion": r.URL.Query().Get("muscle_portion"),
        "joint": r.URL.Query().Get("joint"),
        "movement": r.URL.Query().Get("movement"),
    }
    resources, err := self.AmmRepository.GetAll(filters)
    if err == sql.ErrNoRows {
        self.SuccessResponse(w, r, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
}
