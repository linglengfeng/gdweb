package in_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

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
