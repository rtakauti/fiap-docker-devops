package services

import (
	"api/datamodels"
	"api/repositories"
)

// ContactService interface
type ContactService interface {
	GetAll() []datamodels.Contact
	GetById(id int) (datamodels.Contact, bool)
	Delete(id int) bool
	Insert(item datamodels.Contact) (datamodels.Contact, error)
	Update(id int, item datamodels.Contact) (datamodels.Contact, error)
}

// NewContactService contructor
func NewContactService(repo repositories.ContactRepository) *ContactServiceRepo {
	return &ContactServiceRepo{
		Repo: repo,
	}
}

// ContactService
type ContactServiceRepo struct {
	Repo repositories.ContactRepository
}

// GetById return item by id
func (s *ContactServiceRepo) GetById(id int) datamodels.Contact {
	return s.Repo.GetById(id)
}

// GetAll returns all status.
func (s *ContactServiceRepo) GetAll() []datamodels.Contact {
	return s.Repo.GetAll()
}

// Delete delete item
func (s *ContactServiceRepo) Delete(id int) (deleted bool) {
	return s.Repo.Delete(id)
}

// Insert insert item
func (s *ContactServiceRepo) Insert(item datamodels.Contact) (datamodels.Contact, error) {
	return s.Repo.Insert(item)
}

// Insert insert item
func (s *ContactServiceRepo) Update(id int, item datamodels.Contact) (datamodels.Contact, error) {
	return s.Repo.Update(id, item)
}
