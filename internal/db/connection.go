package db

import (
	"database/sql"
	"fmt"
	"github.com/mark-marushak/bot-english-book/config"
	"github.com/mark-marushak/bot-english-book/logger"
)

func GetPostgresConfig() string {
	var (
		dbUser     string
		dbPassword string
		dbName     string
		dbPort     string
	)

	if err := config.Get().Unmarshal("db.user", &dbUser); err != nil {
		logger.Get().Error("error while getting db.user: %v", err)
	}

	if err := config.Get().Unmarshal("db.password", &dbPassword); err != nil {
		logger.Get().Error("error while getting db.password: %v", err)
	}

	if err := config.Get().Unmarshal("db.name", &dbName); err != nil {
		logger.Get().Error("error while getting db.name: %v", err)
	}

	if err := config.Get().Unmarshal("db.port", &dbPort); err != nil {
		logger.Get().Error("error while getting db.port: %v", err)
	}

	connection := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbPort)
	//connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	//	dbUser, dbPassword, "127.0.0.1", dbPort, dbName)
	return connection
}

var instance *sql.DB

func connect() *sql.DB {
	var err error
	if instance == nil {
		instance, err = sql.Open("postgres", GetPostgresConfig())
		if err != nil {
			logger.Get().Error("Error while connecting to Database: %v", err)
			return nil
		}
	}

	return instance
}
