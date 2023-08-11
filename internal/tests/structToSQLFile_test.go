package tests

import (
	SebbiaDB "github.com/sebbia/DB-todos"
	"testing"
)

type User struct {
	ID   int32 `gorm:"primary_key"`
	Name string
	Age  int
}

type Product struct {
	ID          int `gorm:"primary_key"`
	Description string
	Price       float64
	UserID      int
	User        User `gorm:"foreignKey:product:user_id:user:id"`
}

type Ticket struct {
	ID        int `gorm:"primary_key"`
	Price     float64
	ProductID int
	Product   Product `gorm:"foreignKey:ticket:product_id:product:id"`
}

func TestStructToSQLFile(t *testing.T) {
	db := SebbiaDB.New()

	db.CreateSQLFileMigration("./test_migration/test_00001_init_up.sql", User{}, Product{}, Ticket{})

}
