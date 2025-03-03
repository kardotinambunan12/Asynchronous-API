package config

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	errorhandler "spe_test/error_handler"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	viper.AutomaticEnv()
	viper.GetString("ENVIRONMENT")

	consString := viper.GetString("URL_DATABASE")

	db, err := sql.Open("mysql", consString)

	errorhandler.PanicIfNeeded(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Connecting to database...")
	return db
}
