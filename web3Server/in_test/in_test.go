package in_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"web3Server/proto_go"
	"web3Server/src/gogrpc"

	"golang.org/x/exp/slog"
)

func TestLog(t *testing.T) {
	slog.Info("ssssslog")

	textHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})
	logger := slog.New(textHandler)
	pc, codePath, codeLine, ok := runtime.Caller(1)

	codestr := ""
	funcname := ""
	if !ok {
		codestr = "-"
		funcname = "-"
	} else {
		_, fileName := filepath.Split(codePath)
		codestr = fmt.Sprintf("%s:%d", fileName, codeLine)
		funcname = runtime.FuncForPC(pc).Name()
	}
	fmt.Println(pc, codePath, codeLine, ok, codestr, funcname)

	logger.Debug("debug log test")
	logger.Info("info log test")
	logger.Warn("Warn log test")
	logger.Error("err log test", slog.String("file:", codestr), slog.String("func:", funcname))

}

func TestGogrpc(t *testing.T) {
	fmt.Println("TestGogrpc")
	gogrpc.Start()
	rsp1, err1 := gogrpc.Cli.Say(context.Background(), &proto_go.SayRequest{Name: "aaa"})
	rsp2, err2 := gogrpc.Cli.CommMsg(context.Background(), &proto_go.C2S_Map{Map: map[string]string{"a": "1", "b": "2"}})
	rsp3, err3 := gogrpc.Cli.UserLogincode(context.Background(), &proto_go.C2S_UserLogincode{Account: "1234"})
	rsp4, err4 := gogrpc.Cli.UserLoginauth(context.Background(), &proto_go.C2S_UserLoginauth{Account: "1234", Code: "123456"})

	fmt.Println("TestGogrpc Say:", rsp1, err1)
	fmt.Println("TestGogrpc CommMsg:", rsp2, err2)
	fmt.Println("TestGogrpc UserLogincode:", rsp3, err3)
	fmt.Println("TestGogrpc UserLoginauth:", rsp4, err4)
}
