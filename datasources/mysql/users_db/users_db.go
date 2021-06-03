package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Println("Error loading .env file")
	}

	var (
		host     = os.Getenv("host")
		port     = os.Getenv("port")
		username = os.Getenv("username")
		password = os.Getenv("password")
		schema   = os.Getenv("schema")
	)

	// for postgres: // post = 5432
	// dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
	// "password=%s schema=%s",
	// host, port, username, password, schema)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username, password, host, port, schema,
	)
	var err error
	// for postgres
	// Client, err = sql.Open("postgres", dataSourceName)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully connected !!!")
}
