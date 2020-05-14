package gin_graphql

import (
	"os"
	"sync"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var instance *gorm.DB
var once sync.Once

func Connect(env string) *gorm.DB {
	// Todo: db config file
	mysqlConf := mysql.Config{
		User:                 "root",
		Passwd:               "xxxx",
		Addr:                 "127.0.0.1",
		DBName:               "new_todo",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dsn := mysqlConf.FormatDSN()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetInstance() *gorm.DB {
	env := os.Getenv("ENV")
	once.Do(func() {
		instance = Connect(env)
		if env != "production" {
			instance.LogMode(true)
		}
	})
	return instance
}

func Close() {
	if instance == nil {
		return
	}
	instance.Close()
}
