package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/nnnk7777/go-gin-gorm-mysql/router"
)

var Db *gorm.DB
var err error

func main()  {
	initDb()
	router.InitRouter(Db)
}

func initDb()  {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	PROTOCOL := os.Getenv("MYSQL_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	Db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	Db.Set("gorm:table_options", "ENGINE=InnoDB")
	Migrate(Db)
}

// Migrate DBの初期化（既にあるものは上書きされない）
func Migrate(Db *gorm.DB) {
	// Db.AutoMigrate(&model.User{}, &model.Playlist{}, &model.Reaction{}, &model.Song{})
	Db.AutoMigrate()
}
