package data

import (
	"log"
	"testing"
)

// test data
var users = []User{
	{
		Name:     "Peter Jones",
		Email:    "peter@gmail.com",
		Password: "peter_pass",
	},
	{
		Name:     "John Smith",
		Email:    "john@gmail.com",
		Password: "john_pass",
	},
}

func setup() {
	ThreadDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}

func TestInit(t *testing.T) {
	Init()
	err := Db.Ping()
	if err != nil {
		log.Fatal("ping failed")
	}
	Db.Close()
}
