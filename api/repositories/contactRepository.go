package repositories

import (
	"database/sql"

	"../datamodels"

	_ "github.com/go-sql-driver/mysql"
)

// ContactRepository
type ContactRepository interface {
	GetAll() (results []datamodels.Contact)
	Get(contact datamodels.Contact) (results []datamodels.Contact)
	GetById(id int) (result datamodels.Contact)
	Insert(contact datamodels.Contact) (insertedContact datamodels.Contact, err error)
	Update(id int, contact datamodels.Contact) (updatedContact datamodels.Contact, err error)
	Delete(id int) (deleted bool)
}

type dbContactRepository struct {
	Conn *sql.DB
}

// NewContactRepository contrutor da classe
func NewContactRepository(Conn *sql.DB) *dbContactRepository {
	return &dbContactRepository{
		Conn: Conn,
	}
}

func (repository *dbContactRepository) GetAll() (results []datamodels.Contact) {
	selDB, err := repository.Conn.Query("SELECT * FROM tblContato")

	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		contact := datamodels.Contact{}
		err = selDB.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone, &contact.BirthDate)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, contact)
	}
	return
}

func (repository *dbContactRepository) Get(contact datamodels.Contact) (results []datamodels.Contact) {

	return results
}

func (repository *dbContactRepository) GetById(id int) (result datamodels.Contact) {
	stmt, err := repository.Conn.Prepare("SELECT * FROM tblContato where Id = ?")

	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Query(id)
	for res.Next() {
		contact := datamodels.Contact{}
		err = res.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone, &contact.BirthDate)
		if err != nil {
			panic(err.Error())
		}
		result = contact
	}

	defer res.Close()
	return

}

func (repository *dbContactRepository) Insert(contact datamodels.Contact) (insertedContact datamodels.Contact, err error) {

	return insertedContact, err
}

func (repository *dbContactRepository) Update(id int, contact datamodels.Contact) (updatedContact datamodels.Contact, err error) {

	return updatedContact, err
}

func (repository *dbContactRepository) Delete(id int) (deleted bool) {
	stmt, err := repository.Conn.Prepare("DELETE FROM tblContato where Id = ?")

	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	rows, err := res.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	if rows > 0 {
		deleted = true
	} else {
		deleted = false
	}

	return
}
