package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Repo struct {
	db *sql.DB
}

func New() Repo {
	return Repo{db: CreateConnection()}
}

// create connection with postgres db
func CreateConnection() *sql.DB {
	// load .env file
	err := godotenv.Load("../.env")
	if err != nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Open the connection
	DB, err := sql.Open("mysql", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = DB.Ping()

	if err != nil {
		panic(err)
	}
	return DB
}
