package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func config() {
	d, err := gorm.Open("mysql", "nkindi:nkindi22/simplerest?charset=utf8parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
