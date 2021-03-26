package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/go-gin-gorm-mysql/service"
)

// UserList GET /users したときの動作
func (pc Controller) UserList(c *gin.Context) {
	var s service.UserService
	db := c.MustGet("db").(*gorm.DB)
	p, err := s.GetUserList(db)

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, p)
	}

}
