package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)
type ExerciseRepository struct {
    Db *sql.DB
}
func (self ExerciseRepository) GetExercises() ([]core.Exercise, error) {
    query := `SELECT id, name, description, info_link FROM exercise`
    rows, err := self.Db.Query(query)
    var resources []core.Exercise
    if err != nil {
        if err == sql.ErrNoRows {
            return resources, nil
        }
        return resources, err
    }
    for rows.Next() {
        var resource core.Exercise
        if err := rows.Scan(&resource.Id, &resource.Name, &resource.Description, &resource.InfoLink); err != nil {
            return nil, err
        }
        resources = append(resources, resource)
    }
    return resources, nil
}
func (self ExerciseRepository) GetExerciseDetails(id int64) ([]core.Amm, error) {
    query := `SELECT ea.role, m.name movement, mg.name muscle_group,  mp.name muscle_portion, a.name joint FROM exercise_amm ea
    INNER JOIN articulation_movement_muscle amm on amm.id = ea.amm_id
    INNER JOIN movements m on m.id = amm.movement_id
    INNER JOIN muscle_portion mp on mp.id = amm.muscle_portion_id
    INNER JOIN muscle_group mg on mg.id = mp.muscle_group_id
    INNER JOIN articulations a on a.id = amm.articulation_id
    WHERE ea.exercise_id = ?
    ORDER BY ea.role`
    rows, err := self.Db.Query(query, id)
    var resources []core.Amm
    if err != nil {
        return resources, err
    }
    for rows.Next() {
        var resource core.Amm
        if err := rows.Scan(&resource.Role, &resource.Movement, &resource.MuscleGroup, &resource.MusclePortion, &resource.Joint); err != nil {
            return nil, err
        }
        resources = append(resources, resource)
    }
    return resources, nil
}
