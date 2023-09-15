package logger

import (
	"fmt"

	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"golang.org/x/exp/slog"
)

const (
	Io_way_file     = 1
	Io_way_ctl      = 2
	Io_way_file_ctl = 3
	Nocaller        = -1

	prefix_debug        = "log_debug_"
	prefix_info         = "log_info_"
	prefix_warn         = "log_warn_"
	prefix_error        = "log_error_"
	debug_lv            = "debug"
	info_lv             = "info"
	warn_lv             = "warn"
	error_lv            = "error"
	app_log_lv          = "debug"
	default_caller      = 2
	default_delfileday  = 300
	default_logfilepath = "./logs"
	default_loglv       = "debug"
)

type Logger struct {
	level    string
	self_lv  string
	path     string
	filename string
	time     time.Time
	fd       *os.File
	logger   *slog.Logger
}

var (
	DEBUG         Logger
	INFO          Logger
	WARN          Logger
	ERROR         Logger
	filelock      sync.RWMutex
	levelfilelist = [4]string{prefix_debug, prefix_info, prefix_warn, prefix_error}
	LogCfgValue   LogCfg
)

type formatLogInfo struct {
	msg      string
	file     string
	funcname string
}

type LogCfg struct {
	Loglv      string
	Remain_day int
	Path       string
	Ioway      int
	Caller     int
	Showfile   int
	Showfunc   int
}

func Start(logcfg LogCfg) {
	LogCfgValue = logcfg
	createLogFile()
	loglv := LogCfgValue.Loglv
	logfilepath := LogCfgValue.Path
	DEBUG = NewLogger(loglv, debug_lv, logfilepath)
	INFO = NewLogger(loglv, info_lv, logfilepath)
	WARN = NewLogger(loglv, warn_lv, logfilepath)
	ERROR = NewLogger(loglv, error_lv, logfilepath)
	go logJob()
}

func NewLogger(level string, self_lv string, path string) Logger {
	now := time.Now()
	postFix := now.Format("20060102")
	var prefix string
	switch self_lv {
	case debug_lv:
		prefix = prefix_debug
	case info_lv:
		prefix = prefix_info
	case warn_lv:
		prefix = prefix_warn
	case error_lv:
		prefix = prefix_error
	default:
		prefix = prefix_debug
	}
	filename := logfilename(LogCfgValue.Path, prefix, postFix)

	l := Logger{time: time.Now(), level: level, path: path, filename: filename, self_lv: self_lv}
	f, openfileerr := os.OpenFile(l.filename, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	if openfileerr != nil {
		panic(fmt.Errorf("error opening file: %v", openfileerr.Error()))
	}

	var logger *slog.Logger
	var iow io.Writer
	ioway := LogCfgValue.Ioway
	if ioway == Io_way_file {
		iow = f
	}
	if ioway == Io_way_ctl {
		iow = os.Stdout
	}
	if ioway == Io_way_file_ctl {
		iow = io.MultiWriter(os.Stdout, f)
	}

	switch level {
	case debug_lv:
		logger = slog.New(slog.NewTextHandler(iow, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case info_lv:
		logger = slog.New(slog.NewTextHandler(iow, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case warn_lv:
		logger = slog.New(slog.NewTextHandler(iow, &slog.HandlerOptions{Level: slog.LevelWarn}))
	case error_lv:
		logger = slog.New(slog.NewTextHandler(iow, &slog.HandlerOptions{Level: slog.LevelError}))
	default:
		logger = slog.New(slog.NewTextHandler(iow, &slog.HandlerOptions{Level: slog.LevelDebug}))
	}
	l.logger = logger
	l.fd = f
	return l
}

func GetLogger(loglv string) Logger {
	var loggerinfo Logger
	switch loglv {
	case debug_lv:
		loggerinfo = DEBUG
	case info_lv:
		loggerinfo = INFO
	case warn_lv:
		loggerinfo = WARN
	case error_lv:
		loggerinfo = ERROR
	default:
		loggerinfo = DEBUG
	}
	if !isOneDay(loggerinfo.time) {
		loggerinfo = NewLogger(loggerinfo.level, loggerinfo.self_lv, loggerinfo.path)
	}
	return loggerinfo
}

func Debug(format string, a ...any) {
	formatloginfo := Format(format, a...)
	loggerinfo := GetLogger(debug_lv)
	loggerinfo.logger.Debug(formatloginfo.msg, slog.String("file", formatloginfo.file), slog.String("func", formatloginfo.funcname))
}

func Info(format string, a ...any) {
	formatloginfo := Format(format, a...)
	loggerinfo := GetLogger(info_lv)
	loggerinfo.logger.Info(formatloginfo.msg, slog.String("file", formatloginfo.file), slog.String("func", formatloginfo.funcname))
}

func Warn(format string, a ...any) {
	formatloginfo := Format(format, a...)
	loggerinfo := GetLogger(warn_lv)
	loggerinfo.logger.Warn(formatloginfo.msg, slog.String("file", formatloginfo.file), slog.String("func", formatloginfo.funcname))
}

func Error(format string, a ...any) {
	formatloginfo := Format(format, a...)
	loggerinfo := GetLogger(error_lv)
	loggerinfo.logger.Error(formatloginfo.msg, slog.String("file", formatloginfo.file), slog.String("func", formatloginfo.funcname))
}

func Format(format string, a ...any) formatLogInfo {
	msgstr := fmt.Sprintf(format, a...)
	codestr := "-"
	funcname := "-"
	caller := LogCfgValue.Caller
	if caller != Nocaller && caller > 0 {
		pc, codePath, codeLine, ok := runtime.Caller(caller)
		if ok {
			_, fileName := filepath.Split(codePath)
			if LogCfgValue.Showfile == 0 {
				codestr = "-"
			} else {
				codestr = fmt.Sprintf("%s:%d", fileName, codeLine)
			}
			if LogCfgValue.Showfunc == 0 {
				funcname = "-"
			} else {
				funcname = runtime.FuncForPC(pc).Name()
			}
		}
	}
	return formatLogInfo{msg: msgstr, file: codestr, funcname: funcname}
}

func logJob() {
	c := cron.New(cron.WithSeconds())
	spec := "@daily" // "@daily"  //"*/5 * * * * *"
	c.AddFunc(spec, func() {
		Warn("执行log定时任务。。。")
		now := time.Now()
		createLogFile()
		for i := 0; i < len(levelfilelist); i++ {
			fileprefix := levelfilelist[i]
			closeYesterdayLogFile := fmt.Sprintf(fileprefix+"%s.log", now.Add(-24*time.Hour).Format("20060102"))
			file, _ := os.Open(closeYesterdayLogFile)
			file.Sync()
			file.Close()

			// 删除n天前的日志
			delfileday := LogCfgValue.Remain_day
			removeLogFile := fmt.Sprintf(fileprefix+"%s.log", time.Now().Add(time.Duration(delfileday)*-24*time.Hour).Format("20060102"))
			_, err := os.Open(removeLogFile)
			if err != nil {
				Error("logJob remove file:" + err.Error())
				return
			}
		}

		// go func() {
		// 	// 设置for select 的原因是文件虽然被关闭了，但文件所占的process还在进行中，每10秒轮询一次，执行删除操作，确保文件有被删除
		// loop:
		// 	for {
		// 		select {
		// 		case <-time.After(10 * time.Second):
		// 			removeErr := os.Remove(removeLogFile)
		// 			if removeErr != nil {
		// 				Error(removeErr.Error())
		// 			} else {
		// 				Info("删除日志成功：%s", removeLogFile)
		// 				break loop
		// 			}
		// 			break loop
		// 		}
		// 	}
		// }()

	})
	c.Start()
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func isDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func createLogFile() {
	filelock.Lock()
	defer filelock.Unlock()
	now := time.Now()
	postFix := now.Format("20060102")
	logfilepath := LogCfgValue.Path
	if !isDir(logfilepath) {
		err := os.Mkdir(logfilepath, 0755)
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < len(levelfilelist); i++ {
		fileprefix := levelfilelist[i]
		logFile := logfilename(logfilepath, fileprefix, postFix)
		_, openfileerr := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
		if openfileerr != nil {
			panic(fmt.Errorf("error createLogFile file: %v, error:%v", logFile, openfileerr.Error()))
		}
	}
}

func logfilename(logfilepath string, fileprefix string, postFix string) string {
	return logfilepath + "/" + fileprefix + postFix + ".log"
}

func isOneDay(oldtime time.Time) bool {
	now := time.Now()
	return now.Day() == oldtime.Day() && now.Month() == oldtime.Month() && now.Year() == oldtime.Year()
}
