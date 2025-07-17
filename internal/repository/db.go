package db

import (
	"database/sql"
	//匿名导入
	_ "github.com/go-sql-driver/mysql"

	log "hope_blog/pkg/logger"
)

// 连接数据库,return db 和 nil
func ConnectDB() (*sql.DB, error) {
	//默认连接本地3306端口的mysql,用户名hope,密码123456,数据库名hope_blog
	db_conn := "hope:123456@tcp(localhost:3306)/hope_blog"
	//连接数据库池，并不会检验密码是否对错(用来解析数据库连接)
	db, err := sql.Open("mysql", db_conn)
	if err != nil {
		log.Debug("Failed to connect to database")
		return nil, err
	}
	//连接数据库
	err = db.Ping()
	if err != nil {
		log.Debug("Failed to connect to database")
		return nil, err
	}
	return db, nil
}

// 验证用户登录
func ValidateUser(username, password string) bool {
	// 查询用户
	query := `SELECT password FROM users WHERE username = ?`
	var dbPassword string
	err := db.QueryRow(query, username).Scan(&dbPassword)
	if err != nil {
		return false
	}

	// 比较密码
	if dbPassword == password {
		return true
	}

	return false
}
