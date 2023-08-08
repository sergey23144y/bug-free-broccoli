package tests

import (
	SebbiaDB "github.com/sebbia/DB-todos"
	"testing"
)

type User struct {
	Name string
	Age  int
}

type Product struct {
	UserID      User
	Name        string
	Description string
	Price       float64
}

func TestStructToSQLFile(t *testing.T) {
	db := SebbiaDB.New()

	db.CreateSQLFileMigration("./test_migration/test_00001_init_up.sql", User{}, Product{})

}
