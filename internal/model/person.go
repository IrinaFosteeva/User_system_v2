package model

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PersonUpdate struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}
type Storage interface {
	CreatePerson(p *Person) error
	GetAllPersons() ([]*Person, error)
	GetPersonByID(id int) (*Person, error)
	UpdatePerson(id int, update *PersonUpdate) (*Person, error)
	DeletePerson(id int) error
}
