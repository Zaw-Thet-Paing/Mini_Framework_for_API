package databases

import "os"

type dbconf struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	Db_name  string
}

// DB_DRIVER=mysql
// DB_HOST=localhost
// DB_PORT=3306
// DB_USER=root
// DB_PASSWORD=password
// DB_NAME=ncc_blog
func config() dbconf {
	_db := dbconf{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Db_name:  os.Getenv("DB_NAME"),
	}

	return _db
}
