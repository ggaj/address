package database

import (
	"log"

	"github.com/ggaj/address/src/api/config"
	"github.com/ggaj/address/src/domain/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connect() *gorm.DB {

	db, err := gorm.Open(config.C.Database.DbType, config.C.Database.Dsn)
	if err != nil {
		log.Fatalln(err)
	}

	if config.C.Database.Debug {
		db.LogMode(true)
	}

	if config.C.Database.AutoMigrate {
		db.AutoMigrate(&entities.Address{})
		db.Model(entities.Address{})
	}

	return db
}
