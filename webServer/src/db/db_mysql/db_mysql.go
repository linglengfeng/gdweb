package db_mysql

import (
	"webServer/config"
	"webServer/pkg/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Start() {
	mysqlip := config.Config.GetString("mysql.ip")
	if mysqlip != "" {
		mysqlink := mysql.Link{
			User:     config.Config.GetString("mysql.user"),
			Password: config.Config.GetString("mysql.password"),
			Ip:       mysqlip,
			Port:     config.Config.GetString("mysql.port"),
			Db:       config.Config.GetString("mysql.db"),
		}
		DB = mysql.Start(mysqlink)
	}
}

func GetUseridByAccount(account string) string {
	var userid string
	query := "select id from user where account = ?"
	DB.Raw(query, account).Scan(&userid)
	return userid
}

func InsertUser(account string) string {
	var userid string
	query := "call sp_user_insert(?)"
	DB.Raw(query, account).Scan(&userid)
	return userid
}
