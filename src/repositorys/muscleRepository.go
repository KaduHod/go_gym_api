package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)

type MuscleRepository struct {
	Db *sql.DB
}

func (s *MuscleRepository) GetAll() ([]core.MuscleGroup, error) {
	query := "SELECT id, name FROM muscle_group"
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var resources []core.MuscleGroup
	for rows.Next() {
		var resource core.MuscleGroup
		if err := rows.Scan(&resource.Id, &resource.Name); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s *MuscleRepository) GetById(id int) (core.MuscleGroup, error) {
	query := "SELECT id, name FROM muscle_group WHERE id = ?"
	row := s.Db.QueryRow(query, id)

	var resource core.MuscleGroup
	if err := row.Scan(&resource.Id, &resource.Name); err != nil {
		return resource, err
	}
	return resource, nil
}
func (s *MuscleRepository) GetWithPortions() (*[]core.MuscleGroup, error) {
    groups, err := s.GetAll()
    if err != nil {
       return nil, err
    }
    portions, err := s.GetAllPortions()
    if err != nil {
       return nil, err
    }
    var resources []core.MuscleGroup
    for _, group := range groups {
        for _, portion := range portions {
            if *portion.MuscleGroupId == *group.Id {
                portion.MuscleGroupId = nil
                group.Portions = append(group.Portions, portion)

            }
        }
        resources = append(resources, group)
    }
    return &resources, nil
}

func (s MuscleRepository) GetAllPortions() ([]core.MusclePortion, error) {
	query := "SELECT id, name, muscle_group_id FROM muscle_portion"
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []core.MusclePortion
	for rows.Next() {
		var resource core.MusclePortion
		if err := rows.Scan(&resource.Id, &resource.Name, &resource.MuscleGroupId); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s MuscleRepository) GetPortionByGroupId(groupId int) (*core.MusclePortion, error) {
	query := "SELECT id, name, muscle_group_id FROM muscle_portion WHERE muscle_group_id = ?"
	row := s.Db.QueryRow(query, groupId)

	var resource core.MusclePortion
	if err := row.Scan(&resource.Id, &resource.Name, &resource.MuscleGroupId); err != nil {
		return nil, err
	}
	return &resource, nil
}
func (s MuscleRepository) GetPortionById(id int) (*core.MusclePortion, error) {
	query := "SELECT id, name, muscle_group_id FROM muscle_portion WHERE id = ?"
	row := s.Db.QueryRow(query, id)

	var resource core.MusclePortion
	if err := row.Scan(&resource.Id, &resource.Name, &resource.MuscleGroupId); err != nil {
		return nil, err
	}
	return &resource, nil
}

