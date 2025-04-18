package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
	"strings"
)
type AmmRepository struct {
	Db *sql.DB
}

func (s AmmRepository) GetAll(filters map[string]string) ([]core.MuscleMovementInfo, error) {
	query := `SELECT
		mp.name AS muscle_portion_name,
		mg.name AS muscle_group_name,
		m.name AS movement_name,
		a.name AS joint_name
	FROM articulation_movement_muscle amm
	INNER JOIN muscle_portion mp ON amm.muscle_portion_id = mp.id
	INNER JOIN muscle_group mg ON mp.muscle_group_id = mg.id
	INNER JOIN movements m ON amm.movement_id = m.id
	INNER JOIN articulations a ON amm.articulation_id = a.id`

	var queryParts []string
	var args []interface{}

	if muscleGroup := filters["muscle_group"]; muscleGroup != "" {
		queryParts = append(queryParts, "mg.name = ?")
		args = append(args, strings.TrimSpace(muscleGroup))
	}
	if musclePortion := filters["muscle_portion"]; musclePortion != "" {
		queryParts = append(queryParts, "mp.name = ?")
		args = append(args, strings.TrimSpace(musclePortion))
	}
	if joint := filters["joint"]; joint != "" {
		queryParts = append(queryParts, "a.name = ?")
		args = append(args, strings.TrimSpace(joint))
	}
	if movement := filters["movement"]; movement != "" {
		queryParts = append(queryParts, "m.name = ?")
		args = append(args, strings.TrimSpace(movement))
	}

	if len(queryParts) > 0 {
		query += " WHERE " + strings.Join(queryParts, " AND ")
	}
	rows, err := s.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resources []core.MuscleMovementInfo
	for rows.Next() {
		var resource core.MuscleMovementInfo
		if err := rows.Scan(
			&resource.MusclePortionName,
			&resource.MuscleGroupName,
			&resource.MovementName,
			&resource.JointName,
		); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}
