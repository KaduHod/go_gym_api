package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
	"strings"
)
type ExerciseRepository struct {
    Db *sql.DB
}
func (self ExerciseRepository) GetExercises(search string) ([]core.Exercise, error) {
	var query string

	if search == "" {
		query = `SELECT id, name, description, info_link FROM exercise`
	} else {
		column := "full_text"
		if strings.Contains(search, "structure:") {
			search = strings.Replace(search, "structure:", "", 1)
			column = "structures_text"
		}
		query = `
			SELECT e.id, e.name, e.description, e.info_link
			FROM exercise_search es
			INNER JOIN exercise e ON e.id = es.exercise_id
			WHERE MATCH(es.full_text) AGAINST(? IN NATURAL LANGUAGE MODE)
			ORDER BY MATCH(es.full_text) AGAINST(? IN NATURAL LANGUAGE MODE) DESC
			LIMIT 5
		`
		query = strings.ReplaceAll(query, "full_text", column)
	}
    rows, err := self.Db.Query(query, search, search)
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
func (self ExerciseRepository) GetExercise(id int64) (core.Exercise, error) {
    query := `SELECT id, name, description, info_link FROM exercise WHERE id = ?`
    var resource core.Exercise
    if err := self.Db.QueryRow(query, id).Scan(&resource.Id, &resource.Name, &resource.Description, &resource.InfoLink); err != nil {
        if err == sql.ErrNoRows {
            return resource, nil
        }
        return resource, err
    }
    return resource, nil
}
