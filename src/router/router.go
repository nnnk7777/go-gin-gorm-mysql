package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/go-gin-gorm-mysql/controller"
)

// InitRouter Ginとルーティングの設定
func InitRouter(Db *gorm.DB) {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", Db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!!")
	})
	u := r.Group("/users")
	{
		ctrl := controller.Controller{}
		u.GET("", ctrl.UserList)
	}

	r.Run(":8080")
}
