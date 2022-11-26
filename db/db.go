package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDatabase() {
	if os.Getenv("BE_MODE") == "TEST" {
		dotenvErr := godotenv.Load(".env")
		if dotenvErr != nil {
			log.Println("Error loading .env.test file", dotenvErr.Error())
		}
	} else {
		dotenvERR := godotenv.Load()
		if dotenvERR != nil {
			log.Println("Error loading .env file")
		}
	}
	log.Println("DSN:",os.Getenv("DSN"))
	migrateConnection, err := migrate.New("file://db/migrate", os.Getenv("DSN"))
	if err != nil {
		fmt.Println(err)
		return
	}
	version, _, _ := migrateConnection.Version()
	fmt.Println(version)
	if version != 1 {
		migrateConnection.Migrate(1)
	}
	migrateConnection.Close()
	dbConnection, err := gorm.Open("postgres", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("dbconn: %s", err.Error())
	}
	db = dbConnection
	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}
