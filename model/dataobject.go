package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

//Person is a struct to store personal information
type Person struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
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
