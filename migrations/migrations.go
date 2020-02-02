package migrations

import (
	"gin-example/models"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func RunMigration() {
	path, _ := os.Getwd()
	driver, _ := mysql.WithInstance(models.Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://"+path+"/migrations",
		"mysql",
		driver,
	)
	m.Up()
}
