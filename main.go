package main

import (
	"blog/pkg/setting"
	"blog/routers"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
	"net/http"
)

func main() {
	router := routers.InitRouter()
	/*router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})*/
	//s := &http.Server{} 是 Go 语言标准库中的一个 HTTP 服务结构体，用于创建和管理 HTTP 服务。
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
