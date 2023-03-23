package app

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hafidz98/be_rumbuk_api/helper"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	helper.Info.Println("Connecting to database...")

	err := godotenv.Load(".env")
	helper.PanicIfError(err)

	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	username := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	//DSN format "root@tcp(localhost:3306)/todos_api_db"
	dsn := (username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname)
	helper.Info.Println("Try connect to " + dsn)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		helper.Error.Println("Database connection error")
		panic(err.Error())
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	helper.Info.Println("Database connection established")
	return db
}
