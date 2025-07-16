package main

import (
	"hope_blog/internal/repository"
	"hope_blog/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	_, err := repository.InitDB()
	if err != nil {
		logger.Log.Fatal("Failed to initialize database: ", err)
	}
	//初始化路由
	router := gin.Default()
	router.LoadHTMLGlob("../templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.Run(":8080")
}
