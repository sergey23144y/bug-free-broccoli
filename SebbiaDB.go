package test

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type sebbiaDB interface {
	CustomAutoMigrate(dst interface{}) error
	Connect(Host, Port, Username, Password, DBName, SSLMode string) error // Метод который создает подключение к бд
	Insert(input interface{}) error                                       // Запрос на внесение данных
	GetAll(dest interface{}) error                                        // Запрос ны вывод всей таблицы
	GetById(dest interface{}, id interface{}) error                       // Запрос ны вывод одного элемента таблицы
	Update(dest interface{}, id interface{}) error                        // Запрос ны изменение одного элеммента таблицы
	Delete(dest interface{}, id interface{}) error                        // Запрос ны удаление одного элемента  таблицы
}

func New() sebbiaDB {
	return &DBGORM{}
}

type DBGORM struct {
	db *gorm.DB
}

func (D *DBGORM) CustomAutoMigrate(dst interface{}) error {

	err := D.db.AutoMigrate(&dst)
	if err != nil {
		return err
	}
	log.Println("Все ок")

	return nil
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// Connect подключение базы данных
func (D *DBGORM) Connect(Host, Port, Username, Password, DBName, SSLMode string) error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // Установите свой log.Logger
		logger.Config{
			SlowThreshold: time.Second, // Порог для медленных запросов
			LogLevel:      logger.Info, // Уровень логирования
			Colorful:      true,        // Включить цветной вывод
		},
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		Host, Port, Username, DBName, Password, SSLMode)
	for {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})

		dbSQL, err := db.DB()
		if err != nil {
			log.Printf("Ошибка соединение с БД: %s", err.Error())
			return err
		}

		err = dbSQL.Ping()
		if err == nil {
			log.Println("Соединение с базой данных установлено!")
			D.db = db
			return nil
		}

		log.Println("Соединение с базой данных не установлено. Повторная проверка через 5 секунд...")
		time.Sleep(3 * time.Second)
	}

}

func (D *DBGORM) Insert(input interface{}) error {
	id := D.db.Create(input)
	return id.Error
}

func (D *DBGORM) GetAll(dest interface{}) error {
	result := D.db.Find(dest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil
}

func (D *DBGORM) GetById(dest interface{}, id interface{}) error {
	result := D.db.Where("ID = ?", id).Find(dest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil
}

func (D *DBGORM) Update(dest interface{}, id interface{}) error {
	result := D.db.Model(dest).Where("ID = ?", id).Updates(dest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil

}

func (D *DBGORM) Delete(dest interface{}, id interface{}) error {
	result := D.db.Model(dest).Where("ID = ?", id).Delete(id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Ошибка 0 полей ")
	}
	return nil
}
