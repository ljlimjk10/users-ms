package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

type (
	DBConnector struct {
		DB      *pg.DB
		Options *DBConnectorOptions
	}

	DBConnectorOptions struct {
		User     string
		DB_Name  string
		Addr     string
		Password string
	}
)

func InitDBConnection() *DBConnector {
	dbConnector := &DBConnector{
		Options: &DBConnectorOptions{
			User:     os.Getenv("DB_USER"),
			DB_Name:  os.Getenv("DB_NAME"),
			Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}

	pgOptions := &pg.Options{
		User:     dbConnector.Options.User,
		Database: dbConnector.Options.DB_Name,
		Addr:     dbConnector.Options.Addr,
		Password: dbConnector.Options.Password,
	}
	fmt.Println("Connecting to database.")
	db := pg.Connect(pgOptions)
	checkConnection(db)
	dbConnector.DB = db
	return dbConnector

}

func checkConnection(db *pg.DB) {
	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("Failed to connect to the database: ", err)
		log.Fatal(err)
	}
	fmt.Println("Connected to database.")
}
