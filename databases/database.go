package databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := ""
	if config().Driver == "mysql" {
		dsn = config().User + ":" + config().Password + "@tcp(" + config().Host + ":" + config().Port + ")/" + config().Db_name

		return gorm.Open(mysql.Open(dsn), &gorm.Config{})

	}

	if config().Driver == "postgres" {
		dsn = "host=" + config().Host + " port=" + config().Port + " user=" + config().User + " password=" + config().Password + " dbname=" + config().Db_name + " sslmode=disable"

		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	return nil, fmt.Errorf("unsupported database driver : %s", config().Driver)

}
