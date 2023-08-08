package test

import (
	"gorm.io/gorm"
	"log"
	"testing"
)

type Model struct {
	gorm.Model
	Name     string `validator:"max=20"`
	Password string `validator:"regex=(^[a-zA-Z]+$)"`
}

func TestDBGORM_Connect(t *testing.T) {

	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})

	if err != nil {
		log.Fatalf("Миграция не прошла")
	}

}

func TestDBGORM_Insert(t *testing.T) {

	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	name := "Sergey"
	password := "123321"
	modelInput := &Model{
		Name:     name,
		Password: password,
	}
	err := db.Insert(modelInput)
	log.Printf("Параметр ID: %d", modelInput.ID)
	if err != nil {
		log.Fatalf("Данные не занисли")
	}
}

func TestDBGORM_GetAll(t *testing.T) {

	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", false)

	var models []Model
	err := db.GetAll(&models)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %d", len(models))
	for _, item := range models {
		log.Println(item.Name, item.Password)
	}
}

func TestDBGORM_GetById(t *testing.T) {
	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	var model Model
	err := db.GetById(&model, 24)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %s", model.Name)

}

func TestDBGORM_Update(t *testing.T) {

	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})

	err = db.Update(&Model{

		Name:     "Ivan",
		Password: "123321",
	}, 24)

	if err != nil {
		log.Fatalf("Данные не обновлены")
	}

	var model Model
	err = db.GetById(&model, 24)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %s", model.Name)

}

func TestDBGORM_Delete(t *testing.T) {
	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})

	err = db.Delete(&Model{}, 3, false)

	if err != nil {
		log.Fatalf("Данные не Удалены")
	}
}

func TestDBGORM_Exec(t *testing.T) {
	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})

	row, err := db.Exec("DELETE FROM \"models\" WHERE \"models\".\"id\" = ?", 6)

	if err != nil {
		log.Fatalf("Данные не Удалены")
	}
	if *row == 0 {
		log.Fatalf("Непроизашло ни одного изменения")
	}
}

func TestDBGORM_ExecGet(t *testing.T) {
	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})
	var model Model
	row, err := db.ExecGet("SELECT * FROM \"models\" WHERE ID = ? AND \"models\".\"deleted_at\" IS NULL", &model, 25)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %s", model.Name)

	if *row == 0 {
		log.Fatalf("Непроизашло ни одного изменения")
	}
}

func TestDBGORM_GetPaginatedResult(t *testing.T) {
	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable", true)
	err := db.CustomAutoMigrate(&Model{})

	var models []Model
	err = db.GetAll(&models)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	result, err := db.GetPaginatedResultFromSlice(models, 1, 10)
	if err != nil {
		log.Fatalf("Пагинация не прошла: %s", err.Error())
	}

	log.Printf("Total: %d", result.total)
	log.Printf("Page: %d", result.page)
	log.Printf("Limit: %d", result.limit)
}
