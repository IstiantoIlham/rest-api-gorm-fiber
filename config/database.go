package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rest-api-gorm-fiber/entities"
)

var Database *gorm.DB
var DatabaseURI = "root@tcp(localhost:3306)/gorm_fiber_restapi?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
	var err error
	Database, err = gorm.Open(mysql.Open(DatabaseURI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	err = Database.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}
	return nil
}
