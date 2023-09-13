package request

import (
	"net/http"
	"web3Server/pkg/jwt"
	"web3Server/pkg/myutil"
	"web3Server/src/db"
	"web3Server/src/log"
	"web3Server/src/sendgrid"

	"github.com/gin-gonic/gin"
)

func handle_user_login_code(c *gin.Context) {
	account := c.PostForm("account")
	code := myutil.RandomString(6)
	log.Debug("handle_user_login_code, account:%v, code:%v", account, code)
	is, err := db.SetUserLoginCode(account, code)
	if !is && err != nil {
		log.Debug("handle_user_login_code, is:%v, err:%v", is, err)
		c.JSON(http.StatusOK, MSGF101)
		return
	}
	errsendcode := sendgrid.SendLoginEmail(account, code)
	if errsendcode != nil {
		log.Debug("handle_user_login_code errsendcode:%v", errsendcode)
		c.JSON(http.StatusOK, MSGF101)
		return
	}

	c.JSON(http.StatusOK, MSG100)
}

func handle_user_login_auth(c *gin.Context) {
	account := c.PostForm("account")
	code := c.PostForm("code")
	if account == "" || code == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}
	userid, errget := db.GetUseridByAccount(account)
	if errget != nil {
		userid = db.InsertUser(account)
	}
	log.Debug("handle_user_login_auth, account:%v, code:%v, userid:%v, err:%v", account, code, userid, errget)
	logincode := db.GetUserLoginCode(account)
	if logincode == "" {
		c.JSON(http.StatusOK, retData(MSGF101, "code expired"))
		return
	}
	if code != logincode {
		c.JSON(http.StatusOK, retData(MSGF101, "code error"))
		return
	}
	db.DelUserLoginCode(account)
	mapinfo := map[string]any{"userid": userid, "account": account}
	token, err := jwt.EncodeJwt(mapinfo)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSGF101, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, token))
}

func handle_user_login_token(c *gin.Context) {
	token := c.PostForm("token")
	if token == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}
	account, err := jwt.DecodeJwt(token)
	log.Debug("handle_user_login_token, account:%v, err:%v", account, err)
	if err != nil {
		c.JSON(http.StatusOK, MSGF101)
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, account))
}
