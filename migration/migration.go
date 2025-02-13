package migration

import (
	"github.com/spurtcms/membership/migration/mysql"
	"github.com/spurtcms/membership/migration/postgres"
	"gorm.io/gorm"
)

func AutoMigration(DB *gorm.DB, dbtype any) {

	if dbtype == "postgres" {

		postgres.MigrateTables(DB) //auto migrate table

	} else if dbtype == "mysql" {

		mysql.MigrateTables(DB) //auto migrate table
	}

}
