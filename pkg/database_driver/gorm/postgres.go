package gorm

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB struct
type DB struct {
	Postgres *gorm.DB
}

var dbConnect = &DB{}

// ConnectToPostgreSQL func
func ConnectToPostgreSQL(host, port, username, pass, dbname string, sslmode bool) (*DB, error) {
	var connectionStr string

	if host == "" && port == "" && dbname == "" {
		return nil, errors.New("cannot estabished the connection")
	}

	if port == "APP_DATABASE_POSTGRES_PORT" {
		port = "5432"
	}

	if sslmode {
		connectionStr = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=require", host, username, pass, dbname, port)
	} else {
		connectionStr = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", host, username, pass, dbname, port)
	}

	dial := postgres.Open(connectionStr)

	var err error
	pg, err := gorm.Open(dial, &gorm.Config{
		DryRun: false,
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	dbConnect.Postgres = pg
	return dbConnect, nil
}

// DisconnectPostgres func
func DisconnectPostgres(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		panic("close db")
	}
	err = sqlDb.Close()
	if err != nil {
		logrus.Error(err)
	}
	logrus.Println("Connected with postgres has closed")
}
