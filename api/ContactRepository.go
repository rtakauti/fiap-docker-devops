package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// ContactRepository
type ContactRepository interface {
	GetAll() (results []Contact)
	Get(contact Contact) (results []Contact)
	GetById(id int) (result Contact)
	Insert(contact Contact) (insertedContact Contact, err error)
	Update(id int, contact Contact) (updatedContact Contact, err error)
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

func (repository *dbContactRepository) GetAll() (results []Contact) {
	selDB, err := repository.Conn.Query("SELECT * FROM tblContato")

	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		contact := Contact{}
		err = selDB.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, contact)
	}
	return
}

func (repository *dbContactRepository) Get(contact Contact) (results []Contact) {

	return results
}

func (repository *dbContactRepository) GetById(id int) (result Contact) {
	stmt, err := repository.Conn.Prepare("SELECT * FROM tblContato where Id = ?")

	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Query(id)
	for res.Next() {
		contact := Contact{}
		err = res.Scan(&contact.Id, &contact.Name, &contact.Email, &contact.Phone)
		if err != nil {
			panic(err.Error())
		}
		result = contact
	}

	defer res.Close()
	return

}

func (repository *dbContactRepository) Insert(contact Contact) (insertedContact Contact, err error) {
	stmt, err := repository.Conn.Prepare("INSERT INTO tblContato (Name, Email, Phone) VALUES(?,?,?)")

	res, err := stmt.Exec(contact.Name, contact.Email, contact.Phone)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	contact.Id = int(id)

	insertedContact = contact

	return insertedContact, err
}

func (repository *dbContactRepository) Update(id int, contact Contact) (updatedContact Contact, err error) {
	stmt, err := repository.Conn.Prepare("UPDATE tblContato SET Nome = ?, Email = ?, Phone = ? WHERE Id = ?")

	_, err = stmt.Exec(contact.Name, contact.Email, contact.Phone, id)
	if err != nil {
		panic(err.Error())
	}

	updatedContact = contact
	updatedContact.Id = id

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
