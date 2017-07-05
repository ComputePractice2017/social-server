package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//Person is a struct to store personal information
type Person struct {
	ID    string `json:"id",gorethink:"id"`
	Name  string `json:"name",gorethink:"name"`
	Email string `json:"email",gorethink:"email"`
}

var session *r.Session

func InitSession() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}
func GetPersons() ([]Person, error) {
	res, err := r.DB("social").Table("social").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Person
	err = res.All(&response)
	if err != nil {
		return nil, err
	}

	return response, nil

}
func NewPerson(p Person) (Person, error) {
	res, err := r.UUID().Run(session)
	if err != nil {
		return p, err
	}
	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return p, err
	}
	p.ID = UUID

	res, err = r.DB("social").Table("social").Insert(p).Run(session)
	if err != nil {
		return p, err
	}

	return p, nil
}

func EditPerson(p Person) error {
	_, err := r.DB("social").Table("social").Get(p.ID).Replace(p).Run(session)
	if err != nil {
		return err
	}
	return nil
}
