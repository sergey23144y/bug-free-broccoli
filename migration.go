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

func (d *DBGORM) Migrate(agrs ...interface{}) {}

func (d *DBGORM) createSQLFileMigration(path string, args ...interface{}) error {
	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("gorm"), sqlize.WithMigrationFolder(""))

	for i := 0; i < len(args); i++ {
		_ = newMigration.FromObjects(args[i])
		foreignKey := d.addForeignKey(args[i])

		f, err := os.Create(path)

		if err != nil {
			return err
		}

		defer f.Close()

		_, err2 := f.WriteString("\n" + newMigration.StringUp() + foreignKey + "\n")

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
		field := valType.Field(i)         // получаем поле, по его номеру
		fieldTag := field.Tag.Get("gorm") // получаем значение тега 'gorm'. Считай, как в hash map

		// проверка, что в значениях тега 'gorm'
		// входит 'foreignKey'
		if strings.Contains(fieldTag, foreignKeyTag) {
			splitTag := strings.Split(fieldTag, separator) // разделяем значение тега на массив

			foreignKey = fmt.Sprintf(
				" ALTER TABLE '%s' ADD FOREIGN KEY (%s) REFERENCES '%s' (%s)",
				splitTag[1], splitTag[2], splitTag[3], splitTag[4])
		}
	}

	return
}
