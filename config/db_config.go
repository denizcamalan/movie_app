package config

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB_Connect() (*gorm.DB,error) {

	db, err := gorm.Open("mysql","root:password@tcp(localhost:3306)/movie_archive?parseTime=true")
	if err != nil {
		log.Println("DB is disconnected.")
		return nil,err
	}
	log.Println("DB is connected.")

	return db,nil
}