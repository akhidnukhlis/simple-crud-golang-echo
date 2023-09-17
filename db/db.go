package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	"github.com/akhidnukhlis/config"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

var db *gorm.DB
var err error

func Init() {
	conf := config.GetConfig()

	//initiate database driver
	if conf.DB_DRIVER == "mysql" {
		dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
		db, err = gorm.Open(conf.DB_DRIVER, dbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", conf.DB_DRIVER)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", conf.DB_DRIVER)
		}
	}
	if conf.DB_DRIVER == "postgres" {
		dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.DB_HOST, conf.DB_PORT, conf.DB_USERNAME, conf.DB_NAME, conf.DB_PASSWORD)
		db, err = gorm.Open(conf.DB_DRIVER, dbURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", conf.DB_DRIVER)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", conf.DB_DRIVER)
		}
	}
}

func CreateCon() *gorm.DB {
	return db
}
