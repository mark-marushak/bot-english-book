package db

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"github.com/mark-marushak/bot-english-book/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *gorm.DB

// user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai
func getPostgresConfig() string {
	var (
		config     = koanf.New(".")
		parser     = yaml.Parser()
		dbUser     string
		dbPassword string
		dbName     string
		dbPort     string
	)

	if err := config.Load(file.Provider("config.yml"), parser); err != nil {
		logger.Get().Error("error loading config: %v", err)
	}

	if err := config.Unmarshal("db.user", &dbUser); err != nil {
		logger.Get().Error("error while getting db.user: %v", err)
	}

	if err := config.Unmarshal("db.password", &dbPassword); err != nil {
		logger.Get().Error("error while getting db.password: %v", err)
	}

	if err := config.Unmarshal("db.name", &dbName); err != nil {
		logger.Get().Error("error while getting db.name: %v", err)
	}

	if err := config.Unmarshal("db.port", &dbPort); err != nil {
		logger.Get().Error("error while getting db.port: %v", err)
	}

	//connection := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	//	dbUser, dbPassword, dbName, dbPort)
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword, "127.0.0.1", dbPort, dbName)
	return connection
}

func DB() *gorm.DB {
	if instance == nil {
		//db, err := gorm.Open(postgres.New(postgres.Config{
		//	DSN:                  getPostgresConfig(),
		//	PreferSimpleProtocol: true,
		//}), &gorm.Config{})

		db, err := gorm.Open(postgres.Open(getPostgresConfig()), &gorm.Config{})

		if err != nil {
			logger.Get().Error("[DB] error while opening database: %v", err)
		}

		instance = db
	}

	return instance
}

func PrepareTable() {
	models := []interface{}{
		&model.User{},
		&model.Book{},
		&model.Language{},
		&model.Word{},
	}

	for i := 0; i < len(models); i++ {
		err := createOrUpdate(models[i])

		if err != nil {
			logger.Get().Error("[DB] while explicit including database: %v", err)
		}
	}
}

func createOrUpdate(structure interface{}) (err error) {
	if DB().Migrator().HasTable(structure) {
		err = DB().Migrator().AutoMigrate(structure)
	} else {
		err = DB().Migrator().CreateTable(structure)
	}

	return err
}
