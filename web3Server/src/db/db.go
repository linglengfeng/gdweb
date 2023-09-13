package db

import (
	"fmt"
	"web3Server/src/db/db_mysql"
	"web3Server/src/db/db_redis"
	"web3Server/src/log"
)

func Start() {
	db_mysql.Start()
	db_redis.Start()
}

func UserIsExist(account string) bool {
	return false
}

func SetUserLoginCode(account, code string) (bool, error) {
	return db_redis.SetUserLoginCode(account, code)
}

func GetUserLoginCode(account string) string {
	return db_redis.GetUserLoginCode(account)
}

func DelUserLoginCode(account string) bool {
	return db_redis.DelUserLoginCode(account)
}

func GetUseridByAccount(account string) (string, error) {
	userid, _ := db_redis.GetUseridByAccount(account)
	if userid == "" {
		userid = db_mysql.GetUseridByAccount(account)
	}
	if userid == "" {
		return userid, fmt.Errorf("GetUseridByAccount failed, accout is %v", account) //errors.New("not find user")
	}
	errredisset := db_redis.SetUseridByAccount(account, userid)
	if errredisset != nil {
		log.Warn("db_redis.SetUseridByAccount, account:%v, userid:%v, err:%v", account, userid, errredisset)
	}
	return userid, nil
}

func InsertUser(account string) string {
	userid := db_mysql.InsertUser(account)
	errredisset := db_redis.SetUseridByAccount(account, userid)
	if errredisset != nil {
		log.Warn("db_redis.SetUseridByAccount, account:%v, userid:%v, err:%v", account, userid, errredisset)
	}
	return userid
}
