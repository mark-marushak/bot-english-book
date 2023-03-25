package db

import (
	"fmt"
	"github.com/mark-marushak/bot-english-book/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormInstance *gorm.DB

func Gorm() *gorm.DB {
	if gormInstance == nil {
		gormInstance, _ = gorm.Open(postgres.New(postgres.Config{
			Conn: connect(),
		}), &gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	}

	return gormInstance
}

func PrepareTable() error {
	models := []interface{}{
		&model.Language{},
		&model.Word{},
		&model.Book{},
		&model.User{},
		&model.BookWord{},
		&model.Education{},
		&model.UserKnowledge{},
	}

	for i := 0; i < len(models); i++ {
		err := createOrUpdate(models[i])

		if err != nil {
			return fmt.Errorf("Error when Create or Update model: %v", err)
		}
	}

	var languages int64
	Gorm().Model(&model.Language{}).Where("code in ('en', 'ua')").Count(&languages)
	if languages <= 0 {
		langModels := model.SetupLangs()
		Gorm().Create(&langModels)
	}

	return nil
}

func createOrUpdate(structure interface{}) (err error) {
	if Gorm().Migrator().HasTable(structure) {
		err = Gorm().AutoMigrate(structure)
	} else {
		err = Gorm().Migrator().CreateTable(structure)
	}

	return err
}
