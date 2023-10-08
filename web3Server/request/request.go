package request

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web3Server/config"
	"web3Server/pkg/crypto"
	"web3Server/pkg/jwt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func Start() {
	ginmod := config.Config.GetString("gin.mod")
	if !(ginmod == gin.ReleaseMode || ginmod == gin.DebugMode || ginmod == gin.TestMode) {
		ginmod = gin.DebugMode
	}
	gin.SetMode(ginmod)
	req := gin.Default()
	if ginmod == gin.DebugMode || ginmod == gin.TestMode {
		pprof.Register(req, DevPprof) // http://localhost:9102/dev/pprof/
		// go tool pprof http://localhost:9102/dev/pprof/profile
	}
	timeoutMs := config.Config.GetInt("gin.timeoutMs")
	req.Use(timeoutMiddleware(timeoutMs))
	request(req)
	ipport := config.Config.GetString("gin.ip") + ":" + config.Config.GetString("gin.port")
	// req.Run(ipport)

	server := &http.Server{
		Addr:           ipport,
		Handler:        req,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go server.ListenAndServe()
	fmt.Println("gin start successed..., ip port:", ipport)
	gracefulExitServer(server)
}

func stop() {

}

func timeoutMiddleware(timeoutMs int) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(time.Duration(timeoutMs)*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, MSGF103)
}

func gracefulExitServer(server *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	sig := <-ch
	fmt.Println("获取到系统信号", sig)
	nowTime := time.Now()
	stop()
	// 设置超时 10s
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("-----exited-----", time.Since(nowTime))
}

func retMsg(ret gin.H, format string, a ...any) gin.H {
	msgstr := fmt.Sprintf(format, a...)
	ret[MESSAGE] = msgstr
	return ret
}

func retData(ret gin.H, data any) gin.H {
	ret[MESSAGE] = data
	return ret
}

func JsonBody(c *gin.Context) (map[string]interface{}, error) {
	// // 原始数据
	// body, errread := io.ReadAll(c.Request.Body)
	// if errread != nil {
	// 	return params, errread
	// }
	// 解析 JSON 数据
	var params map[string]interface{}
	if errjson := c.ShouldBindJSON(&params); errjson != nil {
		return params, errjson
	}
	return params, nil
}

func FormBody(c *gin.Context) (map[string]interface{}, error) {
	// 解析表单数据
	params := make(map[string]interface{})
	if errform := c.Request.ParseForm(); errform != nil {
		return params, errform
	}
	// 获取表单参数
	for key, values := range c.Request.PostForm {
		params[key] = values[0]
	}
	return params, nil
}

func handle_test(c *gin.Context) {
	params, err := FormBody(c)
	if err != nil {
		c.JSON(http.StatusOK, retData(MSGF101, err))
		return
	}
	time.Sleep(3 * time.Second)
	fmt.Println("params:", params)
	c.JSON(http.StatusOK, retData(MSG100, params))
}

func handle_encrypt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}

	infostr, err := crypto.Encrypt(info)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, infostr))
}

func handle_decrypt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}
	infostr, err := crypto.Decrypt(info)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, infostr))
}

func handle_encodejwt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}
	mapinfo := map[string]any{"token": info}
	token, err := jwt.EncodeJwt(mapinfo)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, token))
}

func handle_decodejwt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSGF102)
		return
	}
	tokeninfo, err := jwt.DecodeJwt(info)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, tokeninfo))
}

func handle_tokentest(c *gin.Context) {
	account, isAccount := c.Get(SetAccount)
	if !isAccount {
		c.JSON(http.StatusOK, MSGF101)
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, account))
}
