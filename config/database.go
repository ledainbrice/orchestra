package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db = map[string]*gorm.DB{}
var errdb error

func GetDB(key string) *gorm.DB {
	if _, ok := db[key]; !ok {
		conf := env.database[key]
		log.Println(key)
		log.Println(conf.driver.driverModule())
		log.Println(conf.driver.driverUrl(conf))
		db[key], errdb = gorm.Open(conf.driver.driverModule(), conf.driver.driverUrl(conf))
		db[key].LogMode(true)
		if errdb != nil {
			panic(errdb)
		}
	}
	if err := db[key].DB().Ping(); err != nil {
		db[key] = nil
	}

	return db[key]
}
