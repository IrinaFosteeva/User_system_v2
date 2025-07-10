package storage

import "User_system_v2/internal/model"

type Storage interface {
	CreatePerson(p *model.Person) error
	GetAllPersons() ([]*model.Person, error)
	GetPersonByID(id int) (*model.Person, error)
	UpdatePerson(id int, update *model.PersonUpdate) (*model.Person, error)
	DeletePerson(id int) error
}
