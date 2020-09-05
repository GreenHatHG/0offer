package global

import "github.com/jinzhu/gorm"

var (
	MySQLUser = "root"
	MySQLPwd  = "11"
	MySQLHost = "192.168.153.128:33061"
)

var (
	DB *gorm.DB
)
