package repositories

import (
	"database/sql"

	"../datamodels"
)

type StatusRepository interface {
	GetAll() (results []datamodels.Status)
	Get(status datamodels.Status) (results []datamodels.Status)
	GetById(id int64) (result datamodels.Status)
	Insert(status datamodels.Status) (insertedStatus datamodels.Status, err error)
	Update(id int64, status datamodels.Status) (updatedStatus datamodels.Status, err error)
	Delete(id int64) (deleted bool)
}

type dbStatusRepository struct {
	Conn *sql.DB
}

// NewStatusRepository contrutor da classe
func NewStatusRepository(Conn *sql.DB) *dbStatusRepository {
	return &dbStatusRepository{
		Conn: Conn,
	}
}

func (repository *dbStatusRepository) GetAll() (results []datamodels.Status) {
	selDB, err := repository.Conn.Query("SELECT * FROM tblStatus")

	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		status := datamodels.Status{}
		err = selDB.Scan(&status.ID, &status.Name)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, status)
	}
	return
}

func (repository *dbStatusRepository) Get(status datamodels.Status) (results []datamodels.Status) {

	return results
}

func (repository *dbStatusRepository) GetById(id int64) (result datamodels.Status) {

	return result
}

func (repository *dbStatusRepository) Insert(status datamodels.Status) (insertedStatus datamodels.Status, err error) {

	return insertedStatus, err
}

func (repository *dbStatusRepository) Update(id int64, status datamodels.Status) (updatedStatus datamodels.Status, err error) {

	return updatedStatus, err
}

func (repository *dbStatusRepository) Delete(id int64) (deleted bool) {
	return deleted
}
