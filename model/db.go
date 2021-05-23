package model

import (
	"fmt"
	"log"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// func init() {
//     var err error
//     db, err = gorm.Open("sqlite3", "db/sample.db")
//     if err != nil {
//       panic("failed to connect database")
//     }
//     db.AutoMigrate(&User{})
// }


//connect db
const (
	// データベース
	Dialect = "mysql"

	// ユーザー名
	DBUser = "dbuser"

	// パスワード
	DBPass = "P@ssw0rd"

	// プロトコル
	DBProtocol = "tcp(127.0.0.1:3306)"

	// DB名
	DBName = "mini_blog"

	parseTime = "parseTime=true"

	// タイムゾーン(スラッシュはURLエンコード済みのものを設定)
	loc = "Asia%2FTokyo"
)

//meg_dbへの接続
// func connectGorm() *gorm.DB {

func init(){
	var err error
	connectTemplate := "%s:%s@%s/%s?%s&loc=%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass, DBProtocol, DBName, parseTime, loc)
	db, err = gorm.Open(Dialect, connect)
	if err != nil {
		log.Println(err.Error())
	}
	db.AutoMigrate(&User{})
	// return db
}

