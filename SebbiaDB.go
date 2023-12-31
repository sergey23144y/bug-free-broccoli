package SebbiaDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type PaginatedResult struct {
	total int
	limit int
	page  int
	list  []interface{}
}

type sebbiaDB interface {
	CustomAutoMigrate(dst interface{}) error
	Connect(Host, Port, Username, Password, DBName, SSLMode string, loggering bool) error // Метод который создает подключение к бд
	Insert(input interface{}) error                                                       // Запрос на внесение данных
	GetAll(dest interface{}) error                                                        // Запрос ны вывод всей таблицы
	GetById(dest interface{}, id interface{}) error                                       // Запрос ны вывод одного элемента таблицы
	Update(dest interface{}, id interface{}) error                                        // Запрос ны изменение одного элеммента таблицы
	Delete(dest interface{}, id interface{}, softDelete bool) error
	Exec(query string, value ...interface{}) (*int64, error)
	Migrate(args ...interface{}) error
	CreateSQLFileMigration(path string, args ...interface{}) error
	ExecGet(query string, dest interface{}, value ...interface{}) (*int64, error)
	GetPaginatedResultFromSlice(data interface{}, page, limit int) (*PaginatedResult, error)
	MigrateData(path string) error
}

func New() sebbiaDB {
	return &DBGORM{}
}

type DBGORM struct {
	db     *gorm.DB
	Config Config
}

func (D *DBGORM) CustomAutoMigrate(dst interface{}) error {

	err := D.db.AutoMigrate(&dst)
	if err != nil {
		return err
	}
	log.Println("Все ок")

	return nil
}
func (D *DBGORM) ExecGet(query string, dest interface{}, value ...interface{}) (*int64, error) {

	result := D.db.Raw(query, value).Scan(dest)
	if result.Error != nil {
		return nil, result.Error
	}

	return &result.RowsAffected, nil
}

// TODO: Подумать по поводу value. Т.к если пользователь не указывает value, то появляется ошибка
func (D *DBGORM) Exec(query string, value ...interface{}) (*int64, error) {

	result := D.db.Exec(query)
	if result.Error != nil {
		return nil, result.Error
	}

	return &result.RowsAffected, nil
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
func (D *DBGORM) Connect(Host, Port, Username, Password, DBName, SSLMode string, loggering bool) error {
	var newLogger logger.Interface
	D.Config.Host = Host
	D.Config.Port = Port
	D.Config.Username = Username
	D.Config.Password = Password
	D.Config.DBName = DBName
	if loggering {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // Установите свой log.Logger
			logger.Config{
				SlowThreshold: time.Second, // Порог для медленных запросов
				LogLevel:      logger.Info, // Уровень логирования
				Colorful:      true,        // Включить цветной вывод
			},
		)
	}

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
