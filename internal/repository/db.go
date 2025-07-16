package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"hope_blog/config"
	"hope_blog/pkg/logger"
)

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	// 获取数据库配置
	dbConfig := config.GetDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 测试连接
	if err := TestConnection(db); err != nil {
		return nil, err
	}

	return db, nil
}

// TestConnection 测试数据库连接是否正常
func TestConnection(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		logger.Log.Error("Failed to get database instance: ", err)
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Log.Error("Failed to ping database: ", err)
		return err
	}

	logger.Log.Info("Successfully connected to database!")
	return nil
}
