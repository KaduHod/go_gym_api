package services

import (
	"KaduHod/muscles_api/src/core"
	"database/sql"
)
type MovimentService struct {
	Db *sql.DB
}

func (s *MovimentService) GetAll() ([]core.Moviment, error) {
	query := "SELECT id, name FROM movements"
	rows, err := s.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []core.Moviment
	for rows.Next() {
		var resource core.Moviment
		if err := rows.Scan(&resource.Id, &resource.Name); err != nil {
			return nil, err
		}
		resources = append(resources, resource)
	}
	return resources, nil
}

func (s *MovimentService) GetById(id int) (*core.Moviment, error) {
	query := "SELECT id, name FROM movements WHERE id = ?"
	row := s.Db.QueryRow(query, id)

	var resource core.Moviment
	if err := row.Scan(&resource.Id, &resource.Name); err != nil {
		return nil, err
	}
	return &resource, nil
}
