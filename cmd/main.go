package main

import (
	"database/sql"

	db "hope_blog/internal/repository"

	log "hope_blog/pkg/logger"
)

// 连接数据库
func db_init() *sql.DB {
	db, err := db.ConnectDB()
	if err != nil {
		//数据库连接失败抛异常
		panic("test panic")
	}
	return db
}
func main() {
	log.Info("Starting hope_blog")
	database := db_init()
	defer database.Close()
}
