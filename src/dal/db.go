package db

import (
	"database/sql"

	configuration "weather_project/src/config"
)

func CreateDBConn() (*sql.DB, error) {
	// get configuration
	conf := configuration.GetConfig()

	// create db connection "postgres://user:pass@localhost/bookstore"
	connStr := "postgres://" + conf.DBConfig.DBUser + ":" + conf.DBConfig.DBPassword + "@" + conf.DBConfig.DBHost + "/" + conf.DBConfig.DBName

	// open database connection
	dbConn, err := sql.Open("postgres", connStr)

	return dbConn, err
}
