package config

import (
	log "github.com/siruspen/logrus"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DB_Config() (*gorm.DB,error) {

	db, err := gorm.Open("mysql","root@tcp(host.docker.internal:3306)/movie_archive?charset=utf8mb4&parseTime=True")
	if err != nil {
		log.Info("DB is disconnected.")
		return nil,err
	}
	log.Info("DB is connected.")

	return db,nil
}

func DB_Connect() *gorm.DB{
	db,err := DB_Config()
	if  err !=nil{
		log.Error(err)
		return nil
	}
	return db
}