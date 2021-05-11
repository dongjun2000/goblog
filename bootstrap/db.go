package bootstrap

import (
	"goblog/pkg/model"
	"time"
)

// 初始化数据库和ORM
func SetupDB() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求信息
	sqlDB, _ := db.DB()

	// 设置最大链接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲链接数
	sqlDB.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
}