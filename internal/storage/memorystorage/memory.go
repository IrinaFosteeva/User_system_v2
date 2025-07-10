package memorystorage

import (
	"User_system_v2/internal/model"
	"errors"
	"sync"
)

type MemoryStorage struct {
	data   map[int]*model.Person
	mutex  sync.RWMutex
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:   make(map[int]*model.Person),
		nextID: 1,
	}
}

func (s *MemoryStorage) CreatePerson(p *model.Person) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p.ID = s.nextID
	s.nextID++
	s.data[p.ID] = p
	return nil
}

func (s *MemoryStorage) GetAllPersons() ([]*model.Person, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	result := make([]*model.Person, 0, len(s.data))
	for _, p := range s.data {
		result = append(result, p)
	}
	return result, nil
}

func (s *MemoryStorage) GetPersonByID(id int) (*model.Person, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	p, ok := s.data[id]
	if !ok {
		return nil, nil
	}
	return p, nil
}

func (s *MemoryStorage) UpdatePerson(id int, update *model.PersonUpdate) (*model.Person, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	p, ok := s.data[id]
	if !ok {
		return nil, nil
	}

	if update.Name != nil {
		p.Name = *update.Name
	}
	if update.Age != nil {
		p.Age = *update.Age
	}

	return p, nil
}

func (s *MemoryStorage) DeletePerson(id int) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.data[id]; !ok {
		return errors.New("not found")
	}
	delete(s.data, id)
	return nil
}
