package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

type MusclesController struct {
    Controller
}
type MuscleGroup struct {
    Id *int `json:"id,omitempty"`
    Name string `json:"name"`
}
type MusclePortion struct {
    Id *int `json:"id,omitempty"`
    Name string `json:"name"`
    MuscleGroupId *int `json:"muscle_group_id,omitempty"`
}
func (self MusclesController) ListMuscleGroups(w http.ResponseWriter, r *http.Request) {
    query := "SELECT name FROM muscle_group"
    rows, err := self.Db.Query(query)
    defer rows.Close()
    var resources []MuscleGroup
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    for rows.Next() {
        var resource MuscleGroup
        rows.Scan(&resource.Name)
        resources = append(resources, resource)
    }
    SuccessResponse(w, resources, len(resources))
    return
}
func (self MusclesController) ListMusclePortions(w http.ResponseWriter, r *http.Request) {
    query := "SELECT name FROM muscle_portion"
    rows, err := self.Db.Query(query)
    defer rows.Close()
    var resources []MusclePortion
    if err == sql.ErrNoRows {
        SuccessResponse(w, resources, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    for rows.Next() {
        var resource MusclePortion
        rows.Scan(&resource.Name)
        resources = append(resources, resource)
    }
    SuccessResponse(w, resources, len(resources))
    return
}
func (self MusclesController) ListMuscles(w http.ResponseWriter, r *http.Request) {
    query := "SELECT id, name FROM muscle_group"
    rows, err := self.Db.Query(query)
    defer rows.Close()
    var groups []MuscleGroup
    if err == sql.ErrNoRows {
        SuccessResponse(w, groups, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    for rows.Next() {
        var resource MuscleGroup
        rows.Scan(&resource.Id, &resource.Name)
        groups = append(groups, resource)
    }
    query = "SELECT name, muscle_group_id FROM muscle_portion"
    rows, err = self.Db.Query(query)
    defer rows.Close()
    var portions []MusclePortion
    if err == sql.ErrNoRows {
        SuccessResponse[interface{}](w, nil, 0)
        return
    }
    if err != nil {
        fmt.Print(err)
        InternalServerErrorResponse(w, err)
        return
    }
    for rows.Next() {
        var portion MusclePortion
        rows.Scan(&portion.Name, &portion.MuscleGroupId)
        portions = append(portions, portion)
    }
    var resources []interface{}
    for _, group := range groups {
        groupResources := []interface{}{}
        for _, portion := range portions {
            if *portion.MuscleGroupId != *group.Id {
                continue
            }
            portion.MuscleGroupId = nil
            groupResources = append(groupResources, portion)
        }
        resources = append(resources, map[string]interface{}{
            "name": group.Name,
            "portions": groupResources,
        })
    }
    SuccessResponse(w, resources, len(resources))
    return
}
