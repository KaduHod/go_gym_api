package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
)

type MusculoSkeletalController struct {
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
func (self MusculoSkeletalController) ListMuscleGroups(w http.ResponseWriter, r *http.Request) {
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
func (self MusculoSkeletalController) ListMusclePortions(w http.ResponseWriter, r *http.Request) {
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
func (self MusculoSkeletalController) ListMuscles(w http.ResponseWriter, r *http.Request) {
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
type Joint struct {
    Id int `json:"id,omitempty"`
    Name string `json:"name"`
}
func (self MusculoSkeletalController) ListJoints(w http.ResponseWriter, r *http.Request) {
    query := "SELECT name FROM articulations"
    rows, err := self.Db.Query(query)
    defer rows.Close()
    var resources []Joint
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
        var resource Joint
        rows.Scan(&resource.Name)
        resources = append(resources, resource)
    }
    SuccessResponse(w, resources, len(resources))
}
type Moviment struct {
    Id int `json:"id,omitempty"`
    Name string `json:"name"`
}
func (self MusculoSkeletalController) ListMoviments(w http.ResponseWriter, r *http.Request) {
    query := "SELECT name FROM movements"
    rows, err := self.Db.Query(query)
    defer rows.Close()
    var resources []Moviment
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
        var resource Moviment
        rows.Scan(&resource.Name)
        resources = append(resources, resource)
    }
    SuccessResponse(w, resources, len(resources))
}
type MuscleMovementInfo struct {
	MuscleGroupName   string `json:"muscle_group_name,omitempty"`
	MusclePortionName string `json:"muscle_portion_name,omitempty"`
	MovementName      string `json:"movement_name,omitempty"`
	JointName         string `json:"joint_name,omitempty"`
}
func (self MusculoSkeletalController) ListAmm(w http.ResponseWriter, r *http.Request) {
    query := `SELECT
  mp.name AS muscle_portion_name,
  mg.name AS muscle_group_name,
  m.name AS movement_name,
  a.name AS joint_name
FROM articulation_movement_muscle amm
INNER JOIN muscle_portion mp ON amm.muscle_portion_id = mp.id
INNER JOIN muscle_group mg ON mp.muscle_group_id = mg.id
INNER JOIN movements m ON amm.movement_id = m.id
INNER JOIN articulations a ON amm.articulation_id = a.id
`
    //filters
    muscleGroup := r.URL.Query().Get("muscle_group")
    musclePortion := r.URL.Query().Get("muscle_portion")
    joint := r.URL.Query().Get("joint")
    moviment := r.URL.Query().Get("moviment")
    var queryParts []string
    var args []interface{}

    if muscleGroup != "" {
        queryParts = append(queryParts, "mg.name = ?")
        args = append(args, strings.TrimSpace(muscleGroup))
    }
    if musclePortion != "" {
        queryParts = append(queryParts, "mp.name = ?")
        args = append(args, strings.TrimSpace(musclePortion))
    }
    if joint != "" {
        queryParts = append(queryParts, "a.name = ?")
        args = append(args, strings.TrimSpace(joint))
    }
    if moviment != "" {
        queryParts = append(queryParts, "m.name = ?")
        args = append(args, strings.TrimSpace(moviment))
    }

    // Construir a query
    if len(queryParts) > 0 {
        query += " WHERE " + strings.Join(queryParts, " AND ")
    }
    var resources []MuscleMovementInfo
    rows, err := self.Db.Query(query, args...)
    defer rows.Close()
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
        var resource MuscleMovementInfo
        rows.Scan(&resource.MusclePortionName, &resource.MuscleGroupName, &resource.MovementName, &resource.JointName)
        resources = append(resources, resource)
    }
    SuccessResponse(w, resources, len(resources))
}
