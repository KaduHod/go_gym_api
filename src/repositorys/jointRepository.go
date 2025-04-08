package repository

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)

type JointRepository struct {
	Db *sql.DB
}

func (s *JointRepository) GetAll() ([]core.Joint, error) {
	query := "SELECT id, name FROM articulations"
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []core.Joint
	for rows.Next() {
		var resource core.Joint
		if err := rows.Scan(&resource.Id, &resource.Name); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s *JointRepository) GetById(id int) (*core.Joint, error) {
	query := "SELECT id, name FROM articulations WHERE id = ?"
	row := s.Db.QueryRow(query, id)

	var resource core.Joint
	if err := row.Scan(&resource.Id, &resource.Name); err != nil {
		return nil, err
	}
	return &resource, nil
}

