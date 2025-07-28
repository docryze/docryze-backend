package repositories

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB

	DocumentRep DocumentRepository
)

func InitRepository(dsn string) {
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = mysqlDB

	// 自动迁移
	// db.AutoMigrate(&models.User{})

	DocumentRep = DocumentRepository{}
}
