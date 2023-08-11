package tests

import (
	SebbiaDB "github.com/sebbia/DB-todos"
	"testing"
)

type User struct {
	ID   int32 `sql:"primary_key"`
	Name string
	Age  int
}

type Product struct {
	ID          int `sql:"primary_key"`
	Description string
	Price       float64
	UserID      int
	User        User `sql:"foreignKey:product:user_id:user:id"`
}

type Ticket struct {
	ID        int `sql:"primary_key"`
	Price     float64
	ProductID int
	Product   Product `sql:"foreignKey:ticket:product_id:product:id"`
}

func TestStructToSQLFile(t *testing.T) {
	db := SebbiaDB.New()
	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable")
	db.Migrate(User{}, Product{}, Ticket{})
	db.CreateSQLFileMigration("./test_migration/test_00001_init_up.sql", Product{}, Ticket{}, User{})

}
