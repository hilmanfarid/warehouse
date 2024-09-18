package config

import (
	"fmt"
	"golang-warehouse/helper"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DatabaseConnection() *gorm.DB {
	var (
		host     = os.Getenv("MYSQL_HOST")
		port     = os.Getenv("MYSQL_PORT")
		user     = os.Getenv("MYSQL_USER")
		password = os.Getenv("MYSQL_PASSWORD")
		dbName   = os.Getenv("MYSQL_DATABASE")
	)
	fmt.Println(user)
	sqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&parseTime=true", user, password, host, port, dbName)
	fmt.Println(sqlInfo + "&loc=Asia/Jakarta")
	db, err := gorm.Open(mysql.Open(sqlInfo+"&loc=Asia%2FJakarta"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.ErrorPanic(err)

	return db
}
