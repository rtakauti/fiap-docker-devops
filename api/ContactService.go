package main

import "fmt"

// ContactService interface
type ContactService interface {
	GetAll() []Contact
	GetById(id int) (Contact, bool)
	Delete(id int) bool
	Insert(item Contact) (Contact, error)
	Update(id int, item Contact) (Contact, error)
}

// NewContactService contructor
func NewContactService(repo ContactRepository) *ContactServiceRepo {
	return &ContactServiceRepo{
		Repo: repo,
	}
}

// ContactService
type ContactServiceRepo struct {
	Repo ContactRepository
}

// GetById return item by id
func (s *ContactServiceRepo) GetById(id int) Contact {
	return s.Repo.GetById(id)
}

// GetAll returns all status.
func (s *ContactServiceRepo) GetAll() []Contact {
	return s.Repo.GetAll()
}

// Delete delete item
func (s *ContactServiceRepo) Delete(id int) (deleted bool) {
	return s.Repo.Delete(id)
}

// Insert insert item
func (s *ContactServiceRepo) Insert(item Contact) (Contact, error) {
	return s.Repo.Insert(item)
}

// Insert insert item
func (s *ContactServiceRepo) Update(id int, item Contact) (Contact, error) {
	fmt.Println(item)
	return s.Repo.Update(id, item)
}
