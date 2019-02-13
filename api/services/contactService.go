package services

import (
	"../datamodels"
	"../repositories"
)

// ContactService interface
type ContactService interface {
	GetAll() []datamodels.Contact
	GetById(id int64) (datamodels.Contact, bool)
	Delete(id int64) bool
	Insert(item datamodels.Contact) (datamodels.Contact, error)
	Update(id int64, item datamodels.Contact) (datamodels.Contact, error)
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
func (s *ContactServiceRepo) GetById(id int64) datamodels.Contact {
	return s.Repo.GetById(id)
}

// GetAll returns all status.
func (s *ContactServiceRepo) GetAll() []datamodels.Contact {
	return s.Repo.GetAll()
}

// Delete delete item
func (s *ContactServiceRepo) Delete(id int64) (deleted bool) {
	return s.Repo.Delete(id)
}

// Insert insert item
func (s *ContactServiceRepo) Insert(item datamodels.Contact) (datamodels.Contact, error) {
	return s.Repo.Insert(item)
}

// Insert insert item
func (s *ContactServiceRepo) Update(id int64, item datamodels.Contact) (datamodels.Contact, error) {
	return s.Repo.Update(id, item)
}
