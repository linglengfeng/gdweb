package log

import (
	"webServer/config"
	"webServer/pkg/logger"
)

func Start() {
	loglv := config.Config.GetString("log.level")
	remain_day := config.Config.GetInt("log.remain_day")
	path := config.Config.GetString("log.path")
	showfile := config.Config.GetInt("log.showfile")
	showfunc := config.Config.GetInt("log.showfunc")
	caller := 3
	logcfg := logger.LogCfg{
		Loglv:      loglv,
		Remain_day: remain_day,
		Path:       path,
		Ioway:      logger.Io_way_file_ctl,
		Caller:     caller,
		Showfile:   showfile,
		Showfunc:   showfunc,
	}
	logger.Start(logcfg)
}

func Debug(format string, a ...any) {
	logger.Debug(format, a...)
}

func Info(format string, a ...any) {
	logger.Info(format, a...)
}

func Warn(format string, a ...any) {
	logger.Warn(format, a...)
}

func Error(format string, a ...any) {
	logger.Error(format, a...)
}
