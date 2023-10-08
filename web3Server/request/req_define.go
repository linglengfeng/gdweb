package request

import (
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
	req.Use(func(c *gin.Context) {
		// 检查请求是否是 OPTIONS 请求
		if c.Request.Method == "OPTIONS" {
			// 添加允许的 CORS 标头
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type")
			c.Status(200)
			c.Abort()
			return
		}
		is, err := shouldDisableRoute(c)
		if !is {
			log.Info("request can't used, err:%v", err)
			c.JSON(http.StatusOK, retMsg(MSGF102, err))
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})

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
