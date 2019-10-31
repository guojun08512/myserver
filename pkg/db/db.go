package db

import (
	"fmt"

	"myserver/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

const (
	driverMysql = "mysql"
	// DriverSqlite     = "sqlite"
	// DriverPostgresql = "postgresql"
)

//GetDbClient 获取db instance
func GetDbClient() *gorm.DB {
	return db
}

func init() {
	//open a db connection
	var err error
	dburl := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True", config.GetConfig().DB.UserName, config.GetConfig().DB.PassWord, config.GetConfig().DB.Host, config.GetConfig().DB.DBName)
	db, err = gorm.Open(driverMysql, dburl)
	if err != nil {
		panic("failed to connect database")
	}
	//Migrate the schema
	// db.AutoMigrate(&todoModel{})
}
