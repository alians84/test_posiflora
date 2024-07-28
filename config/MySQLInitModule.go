package config

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"test_posiflora/app/models"
)

var Db *gorm.DB
var err error

func initDatabase() bool {
	mysqlUserName, _ := os.LookupEnv("MYSQL_USER")
	mysqlPassword, _ := os.LookupEnv("MYSQL_PASSWORD")
	mysqlHost, _ := os.LookupEnv("MYSQL_HOST")
	mysqlPort, _ := os.LookupEnv("MYSQL_PORT")
	mysqlDatabase, _ := os.LookupEnv("MYSQL_DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db = Db.Debug()

	Db.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatal("DB initialization failed!")
		panic(err)
		return false
	}
	return true

}
