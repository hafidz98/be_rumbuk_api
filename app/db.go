package app

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hafidz98/be_rumbuk_api/helper"
)

func NewDB() *sql.DB {
	helper.Info.Println("Connecting to database...")

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	//DSN format "root@tcp(localhost:3306)/todos_api_db"
	dsn := (username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname)
	helper.Info.Println("Try connect to " + dsn)

	db, err := sql.Open(driver, dsn)
	helper.PanicIfError(err)

	err = db.Ping()
	if err != nil {
		helper.Error.Println("Database connection error")
		helper.PanicIfError(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	helper.Info.Println("Database connection established")
	return db
}
