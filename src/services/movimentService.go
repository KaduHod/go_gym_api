package services

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)
type MovementService struct {
	Db *sql.DB
}

func (s *MovementService) GetAll() ([]core.Movement, error) {
	query := "SELECT id, name FROM movements"
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []core.Movement
	for rows.Next() {
		var resource core.Movement
		if err := rows.Scan(&resource.Id, &resource.Name); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s *MovementService) GetById(id int) (*core.Movement, error) {
	query := "SELECT id, name FROM movements WHERE id = ?"
	row := s.Db.QueryRow(query, id)

	var resource core.Movement
	if err := row.Scan(&resource.Id, &resource.Name); err != nil {
		return nil, err
	}
	return &resource, nil
}
