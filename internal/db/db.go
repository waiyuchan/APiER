package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database: ", err)
	}

	// 可以在这里配置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		log.Panic("Failed to get database connection pool: ", err)
	}
	sqlDB.SetMaxIdleConns(10)   // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)  // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(0) // 设置了连接可复用的最大时间
}

func GetDB() *gorm.DB {
	return DB
}
