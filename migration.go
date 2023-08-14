package SebbiaDB

import (
	"fmt"
	"github.com/sunary/sqlize"
	"os"
	"reflect"
	"strings"
)

const (
	foreignKeyTag = "foreignKey" // необходимо для определения вхождения в значение тега 'gorm'.
	separator     = ":"          // разделитель для значений тега 'gorm'.
)

func (d *DBGORM) Migrate(args ...interface{}) error {
	for i := 0; i < len(args); i++ {
		if err := d.db.AutoMigrate(args[i]); err != nil {
			return err
		}
	}

	return nil
}

func (d *DBGORM) CreateSQLFileMigration(path string, args ...interface{}) error {
	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("sql"), sqlize.WithMigrationFolder(""))
	_ = os.Remove(path)
	_ = newMigration.FromObjects(args...)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	defer f.Close()
	_, err2 := f.WriteString("\n" + newMigration.StringUp() + "\n")

	if err2 != nil {
		return err
	}

	for i := 0; i < len(args); i++ {
		foreignKey := d.addForeignKey(args[i])

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			return err
		}

		defer f.Close()
		_, err2 := f.WriteString("\n" + foreignKey + "\n")

		if err2 != nil {
			return err
		}
	}

	return nil
}

func (d *DBGORM) addForeignKey(table interface{}) (foreignKey string) {
	val := reflect.ValueOf(table) // в val заноситься реализация конректного типа table
	valType := val.Type()         // определение типа перменной val. Переменная valType становится типом данных

	// итерируемся по количеству полей в структуре
	//TODO: паникует, если valType не структура. Надо исправить.
	for i := 0; i < valType.NumField(); i++ {
		field := valType.Field(i)        // получаем поле, по его номеру
		fieldTag := field.Tag.Get("sql") // получаем значение тега 'gorm'. Считай, как в hash map

		// проверка, что в значениях тега 'gorm'
		// входит 'foreignKey'
		if strings.Contains(fieldTag, foreignKeyTag) {
			splitTag := strings.Split(fieldTag, separator) // разделяем значение тега на массив

			foreignKey = fmt.Sprintf(
				"ALTER TABLE '%s' ADD FOREIGN KEY (%s) REFERENCES '%s' (%s);",
				splitTag[1], splitTag[2], splitTag[3], splitTag[4])
		}
	}

	return
}

// MigrateData принимает на вход путь к sql файлу. После чего производит внос данных в базу данных.
func (d *DBGORM) MigrateData(path string) error {
	// чтение файла, по пути, который передал пользователь
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// т.к в конце каждого sql запроса должна находиться ;
	// то будет производить разделение по ней
	// и уже в цикле отправлять по одному запросу.
	// это сделано из-за того, что gorm в одном exec
	// не может отправлять и обрабатывать несколько запросов.
	splitData := strings.Split(string(data), ";")
	for i := 0; i < len(splitData); i++ {
		_, err = d.Exec(splitData[i], nil)
		if err != nil {
			return err
		}
	}

	return nil
}
