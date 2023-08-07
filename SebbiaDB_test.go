package test

import (
	"gorm.io/gorm"
	"log"
	"testing"
)

type Model struct {
	gorm.Model
	Name     string
	Password string
}

func TestDBGORM_Connect(t *testing.T) {

	db := New()

	db.Connect("localhost", "5433", "loyalty", "loyalty", "loyalty", "disable")

	err := db.CustomAutoMigrate(&Model{})

	if err != nil {
		log.Fatalf("Миграция не прошла")
	}

	name := "Sergey"
	password := "123321"
	modelInput := &Model{
		Name:     name,
		Password: password,
	}
	err = db.Insert(modelInput)
	log.Printf("Параметр имя: %d", modelInput.ID)
	if err != nil {
		log.Fatalf("Данные не занисли")
	}

	var models []Model
	err = db.GetAll(&models)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %d", len(models))
	for _, item := range models {
		log.Println(item.Name, item.Password)
	}
	var model Model
	err = db.GetById(&model, 24)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	log.Printf("Параметр имя: %s", model.Name)

	err = db.Update(&Model{

		Name:     "Max",
		Password: password,
	}, 24)

	if err != nil {
		log.Fatalf("Данные не обновлены")
	}

	log.Printf("Параметр имя: %s", model.Name)
	err = db.GetById(&model, 24)

	if err != nil {
		log.Fatalf("Данные не получены")
	}

	err = db.Delete(&model, 22)

	if err != nil {
		log.Fatalf("Данные не Удалены")
	}

}
