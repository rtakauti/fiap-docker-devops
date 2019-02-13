package services

import (
	"../datamodels"
	"../repositories"
)

// StatusService interface
type StatusService interface {
	GetAll() []datamodels.Status
	GetById(id int64) (datamodels.Status, bool)
	Delete(id int64) bool
	Insert(item datamodels.Status) (datamodels.Status, error)
	Update(id int64, item datamodels.Status) (datamodels.Status, error)
}

// NewStatusService contructor
func NewStatusService(repo repositories.StatusRepository) *StatusServiceTop {
	return &StatusServiceTop{
		Repo: repo,
	}
}

type StatusServiceTop struct {
	Repo repositories.StatusRepository
}

// GetById return item by id
func (s *StatusServiceTop) GetById(id int64) datamodels.Status {
	return s.Repo.GetById(id)
}

// GetAll returns all status.
func (s *StatusServiceTop) GetAll() []datamodels.Status {
	return s.Repo.GetAll()
}

// Delete delete item
func (s *StatusServiceTop) Delete(id int64) (deleted bool) {
	return s.Repo.Delete(id)
}

// Insert insert item
func (s *StatusServiceTop) Insert(item datamodels.Status) (datamodels.Status, error) {
	return s.Repo.Insert(item)
}

// Insert insert item
func (s *StatusServiceTop) Update(id int64, item datamodels.Status) (datamodels.Status, error) {
	return s.Repo.Update(id, item)
}
