package request

import (
	"gdback/pkg/myutil"
	"net/http"
	"web3Server/config"
	"web3Server/pkg/jwt"
	"web3Server/src/db"
	"web3Server/src/log"

	"github.com/gin-gonic/gin"
)

const (
	// return msg fileds
	STATE   = "state"
	MESSAGE = "message"

	// other fields
	SetAccount = "account"

	//STATE
	S100     = 100
	S100_MSG = "success"
	F101     = -101
	S101_MSG = "failed"
	F102     = -102
	F102_MSG = "parameter error"
	F103     = -103
	F103_MSG = "timeout"

	// GET
	DevPprof = "/dev/pprof"

	// POST
	Test      = "/test"
	Encrypt   = "/encrypt"
	Decrypt   = "/decrypt"
	EncodeJwt = "/encodeJwt"
	DecodeJwt = "/decodeJwt"
	Tokentest = "/tokentest"

	UserLoginCode  = "/user/login/code"
	UserLoginAuth  = "/user/login/auth"
	UserLoginToken = "/user/login/token"
)

var (
	// limit
	LimitApi           = map[string]byte{Test: 1, Encrypt: 1, Decrypt: 1, EncodeJwt: 1, DecodeJwt: 1, Tokentest: 1}
	UnAuthorizationApi = map[string]byte{Test: 1, Encrypt: 1, Decrypt: 1, UserLoginCode: 1, UserLoginAuth: 1, UserLoginToken: 1}

	MSG100 = gin.H{
		"state":   S100,
		"message": S100_MSG,
	}

	MSGF101 = gin.H{
		"state":   F101,
		"message": S101_MSG,
	}

	MSGF102 = gin.H{
		"state":   F102,
		"message": F102_MSG,
	}

	MSGF103 = gin.H{
		"state":   F103,
		"message": F103_MSG,
	}
)

func request(req *gin.Engine) {
	req.Use(headerMiddleware, allowedIPsMiddleware)

	//get

	//post
	req.POST(Test, handle_test)
	req.POST(Encrypt, handle_encrypt)
	req.POST(Decrypt, handle_decrypt)
	req.POST(Tokentest, handle_tokentest)
	req.POST(EncodeJwt, handle_encodejwt)
	req.POST(DecodeJwt, handle_decodejwt)
	req.POST(UserLoginCode, handle_user_login_code)
	req.POST(UserLoginAuth, handle_user_login_auth)
	req.POST(UserLoginToken, handle_user_login_token)
}

func headerMiddleware(c *gin.Context) {
	// 检查请求是否是 OPTIONS 请求
	if c.Request.Method == "OPTIONS" || c.Request.Method == "POST" || c.Request.Method == "GET" {
		// 允许特定域的跨域请求
		c.Header("Access-Control-Allow-Origin", "*")
		// 允许特定的 HTTP 方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		// 允许特定的请求头
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// 允许跨域请求包含凭据（如 Cookie）
		c.Header("Access-Control-Allow-Credentials", "true")
		// 设置预检请求有效期
		c.Header("Access-Control-Max-Age", "600")
		if c.Request.Method == "OPTIONS" {
			c.Status(200)
			c.Abort()
			return
		}
	}
	is, err := shouldDisableRoute(c)
	if !is {
		log.Info("request can't used, err:%v", err)
		c.JSON(http.StatusOK, retMsg(MSGF102, err))
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
	c.Next()
}

func allowedIPsMiddleware(c *gin.Context) {
	isAllowedIPs := config.Config.GetInt("isAllowedIPs")
	if isAllowedIPs == 1 {
		clientIP := c.ClientIP()
		allowedIPs := config.Config.GetStringSlice("allowedIPs")
		ismem := myutil.IsMember[string](clientIP, allowedIPs)
		if !ismem {
			log.Info("Access denied, clientIP:%v", clientIP)
			c.JSON(http.StatusOK, retMsg(MSGF102, "Access denied"))
			c.Abort()
			return
		}
	}
	c.Next()
}

func shouldDisableRoute(c *gin.Context) (bool, string) {
	fullpath := c.FullPath()
	checkLimitApi, errapi := checkLimitApi(fullpath)
	if !checkLimitApi {
		return checkLimitApi, errapi
	}

	token := c.GetHeader("Authorization")
	checkAuth, errauth, account := checkAuth(fullpath, token)
	if !checkAuth {
		return checkAuth, errauth
	}
	c.Set(SetAccount, account)
	return true, ""
}

func checkLimitApi(fullpath string) (bool, string) {
	serverType := config.Config.GetString("server_type")
	if serverType == "dev" {
		return true, "isLimitApi"
	} else {
		return LimitApi[fullpath] != 1, "isLimitApi"
	}
}

func checkAuth(fullpath, token string) (bool, string, string) {
	retmsg := ""
	account := ""
	if UnAuthorizationApi[fullpath] == 1 {
		return true, retmsg, account
	}
	mapinfo, err := jwt.DecodeJwt(token)
	if err != nil {
		return false, err.Error(), ""
	}
	info := mapinfo["info"].(map[string]any)
	if info["account"] == nil {
		return false, "Authorization error", account
	}
	account = info["account"].(string)
	if !db.UserIsExist(account) {
		return false, "account not exist", account
	}
	return true, retmsg, account
}
