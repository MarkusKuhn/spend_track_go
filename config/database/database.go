package database

import (
	"database/sql"
	"fmt"
	"spend_track_go/helpers/utilities"
)

func OpenConnection(connectionString string) *sql.DB {
	psqlConnectionString := ConnectionString()

	db, err := sql.Open("postgres", psqlConnectionString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}

func ConnectionString() string {
	host := utilities.GetEnvVariable("PG_HOST")
	port := utilities.GetEnvVariable("PG_PORT")
	user := utilities.GetEnvVariable("PG_USER")
	password := utilities.GetEnvVariable("PG_PASSWORD")
	dbname := utilities.GetEnvVariable("PG_DBNAME")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}
