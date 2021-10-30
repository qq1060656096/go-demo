package example

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

func init()  {
	db = Connect()
	db.Where("1=1").Delete(demo{})
}

func Connect() *gorm.DB {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/godemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("mysql connect fail: ", err)
	}
	return db
}


type demo struct {
	Name string
}

func (d demo) TableName() string {
	return "demo"
}
