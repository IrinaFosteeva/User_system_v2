package postgresstorage

import (
	"User_system_v2/internal/model"
	"database/sql"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreatePerson(p *model.Person) error {
	err := s.db.QueryRow(`
		INSERT INTO persons (name, age) VALUES ($1, $2) RETURNING id
	`, p.Name, p.Age).Scan(&p.ID)
	return err
}

func (s *PostgresStorage) GetAllPersons() ([]*model.Person, error) {
	rows, err := s.db.Query(`SELECT id, name, age FROM persons`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []*model.Person
	for rows.Next() {
		p := &model.Person{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Age); err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func (s *PostgresStorage) GetPersonByID(id int) (*model.Person, error) {
	p := &model.Person{}
	err := s.db.QueryRow(`SELECT id, name, age FROM persons WHERE id = $1`, id).
		Scan(&p.ID, &p.Name, &p.Age)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return p, err
}

func (s *PostgresStorage) UpdatePerson(id int, update *model.PersonUpdate) (*model.Person, error) {
	p, err := s.GetPersonByID(id)
	if err != nil || p == nil {
		return nil, err
	}

	if update.Name != nil {
		p.Name = *update.Name
	}
	if update.Age != nil {
		p.Age = *update.Age
	}

	_, err = s.db.Exec(`UPDATE persons SET name=$1, age=$2 WHERE id=$3`, p.Name, p.Age, p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PostgresStorage) DeletePerson(id int) error {
	_, err := s.db.Exec(`DELETE FROM persons WHERE id = $1`, id)
	return err
}
