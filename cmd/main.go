package main

import (
	"database/sql"

	db "hope_blog/internal/repository"

	log "hope_blog/pkg/logger"

	"github.com/gin-gonic/gin"
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

// 渲染登录页面
func login(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"title": "Gin_index"})
}

// 处理登录逻辑
func handleLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user, err := db.ValidateUser(username, password)
	if err != nil {
		c.HTML(200, "index.html", gin.H{"error": "用户名或密码错误"})
	}

	c.HTML(200, "index.html", gin.H{"message": "登录成功"})
	return user
}
func main() {
	log.Info("Starting hope_blog")
	database := db_init()
	defer database.Close()
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/", login)
	router.POST("/login", handleLogin)
	router.Run(":8080")
}
