package logic

import (
	"fmt"
	"web3Server/pkg/jwt"
	"web3Server/pkg/myutil"
	"web3Server/src/db"
	"web3Server/src/log"
	"web3Server/src/sendgrid"
)

func UserLogincode(account string) error {
	code := myutil.RandomString(6)
	is, err := db.SetUserLoginCode(account, code)
	if !is && err != nil {
		return err
	}
	return sendgrid.SendLoginEmail(account, code)
}

func UserLoginauth(account, code string) (string, error) {
	userid, errget := db.GetUseridByAccount(account)
	if errget != nil {
		userid = db.InsertUser(account)
	}
	log.Debug("UserLoginauth, account:%v, code:%v, userid:%v, err:%v", account, code, userid, errget)
	logincode := db.GetUserLoginCode(account)
	if logincode == "" {
		return "", fmt.Errorf("code expired")
	}
	if code != logincode {
		return "", fmt.Errorf("code error")
	}
	db.DelUserLoginCode(account)
	mapinfo := map[string]any{"userid": userid, "account": account}
	return jwt.EncodeJwt(mapinfo)
}
