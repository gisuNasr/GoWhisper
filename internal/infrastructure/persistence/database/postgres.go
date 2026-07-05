package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(DSN string) error {
	var err error

	dbClient, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(25)
	sqlDb.SetMaxOpenConns(25)
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Db connection established")
	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}
