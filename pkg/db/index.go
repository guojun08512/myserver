package db

import (
	"errors"
	"fmt"
	"myserver/pkg/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

const (
	driverMysql = "mysql"
	// DriverSqlite     = "sqlite"
	// DriverPostgresql = "postgresql"
)

var Gorm map[string]*gorm.DB

func init() {
	Gorm = make(map[string]*gorm.DB)
}

// 初始化Gorm
func NewDB(dbname string) {

	var orm *gorm.DB
	var err error

	//默认配置
	dbHost := config.GetConfig().DB.Host
	dbUser := config.GetConfig().DB.UserName
	dbPasswd := config.GetConfig().DB.PassWord
	idleconnsMax := config.GetConfig().DB.DBIdleconnsMax
	openconnsMax := config.GetConfig().DB.DBOpenconnsMax

	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPasswd, dbHost, dbname)
	//开启sql调试模式
	//GDB.LogMode(true)
	retryCount := 5
	for orm, err = gorm.Open(driverMysql, connectString); err != nil; {
		fmt.Println("数据库连接异常! 5秒重试")
		time.Sleep(5 * time.Second)
		orm, err = gorm.Open(driverMysql, connectString)
		if retryCount < 0 {
			panic(errors.New(" 数据库连接失败!! "))
		}
		retryCount--
	}
	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(idleconnsMax)
	//最大打开连接数
	orm.DB().SetMaxIdleConns(openconnsMax)
	Gorm[dbname] = orm
	orm.AutoMigrate(&User{})
}

// GetORMByName 通过名称获取Gorm实例
func GetORMByName(dbname string) *gorm.DB {
	return Gorm[dbname]
}

// GetORM 获取默认的Gorm实例
func GetORM() *gorm.DB {
	return Gorm[config.GetConfig().DB.DBName]
}

// CloseORM 系统关闭退出
func CloseORM() {
	for _, v := range Gorm {
		v.Close()
	}
}
