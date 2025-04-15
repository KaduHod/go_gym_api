package controllers

import (
	repository "KaduHod/muscles_api/src/repositorys"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ExerciseController struct {
    Controller
    ExerciseRepository *repository.ExerciseRepository
}
// ListExercises godoc
// @Summary List all exercises
// @Description Get a list of all exercises with their information
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=[]core.ExerciseSimple,metadata=controllers.MetaData}
// @Success 304 "Not modified"
// @Header 304 {string} ETag "Entity tag for cache validation"
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /exercises [get]
func (self ExerciseController) GetExercises(w http.ResponseWriter, r *http.Request) {
    resources, err := self.ExerciseRepository.GetExercises()
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    self.SuccessResponse(w, r, resources, len(resources))
}
// ListExercises godoc
// @Summary List all exercises movements
// @Description Get a list of all movements on the execution of the exercise
// @Tags Musculoskeletal
// @Accept  json
// @Produce  json
// @Success 200 {object} controllers.ResponseDescription{data=core.Exercise,metadata=controllers.MetaData}
// @Success 304 "Not modified"
// @Header 304 {string} ETag "Entity tag for cache validation"
// @Failure 500 {object} controllers.ResponseDescription{data=string}
// @Security BearerAuth
// @Failure 401 {object} controllers.ResponseUnauthorized
// @Router /exercises/{id} [get]
func (self ExerciseController) GetExercise(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    idInt, err := strconv.ParseInt(id, 10, 64)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    exercise, err := self.ExerciseRepository.GetExercise(idInt)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    if exercise.Id == 0 {
        self.SuccessResponse(w, r, exercise, 0)
        return
    }
    resources, err := self.ExerciseRepository.GetExerciseDetails(idInt)
    if err != nil {
        self.InternalServerError(w, r, err)
        return
    }
    exercise.Map = resources
    self.SuccessResponse(w, r, exercise, len(exercise.Map))
}
