package initialize

import (
	"0offer/global"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func Mysql() {
	str := fmt.Sprintf("%s:%s@(%s)/miaosha?charset=utf8&parseTime=True&loc=Local",
		global.MySQLUser, global.MySQLPwd, global.MySQLHost)
	if db, err := gorm.Open("mysql", str); err != nil {
		log.Fatalf("connect MySQL failed with %s\n", err)
		os.Exit(-1)
	} else {
		global.DB = db
		global.DB.DB().SetMaxIdleConns(10)
		global.DB.DB().SetMaxOpenConns(10)
		global.DB.SingularTable(true)
	}
}
