package SebbiaDB

import (
	"github.com/sunary/sqlize"
	"os"
)

func (d *DBGORM) Migrate(agrs ...interface{}) {}

func (d *DBGORM) CreateSQLFileMigration(path string, args ...interface{}) error {
	newMigration := sqlize.NewSqlize(sqlize.WithSqlTag("psql"), sqlize.WithMigrationFolder(""))

	for i := 0; i < len(args); i++ {
		_ = newMigration.FromObjects(args[i])

		f, err := os.Create(path)

		if err != nil {
			return err
		}

		defer f.Close()

		_, err2 := f.WriteString(newMigration.StringUp() + "\n")

		if err2 != nil {
			return err
		}
	}

	return nil
}
