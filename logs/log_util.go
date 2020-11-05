package logs

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// 日志保留天数：保留30天
const Log_SAVE_DAYS = 30

func NewLfsHook(logger *(logrus.Logger), logName string, logLevel logrus.Level, maxRemainCnt uint) logrus.Hook {
	writer, err := rotatelogs.New(
		//logName+".%Y%m%d%H",
		logName+".%Y%m%d.log",
		//logName+".%Y%m%d%H%M.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(logName),

		// WithRotationTime设置日志分割的时间，这里设置为24小时分割一次
		rotatelogs.WithRotationTime(time.Hour*24),
		//rotatelogs.WithRotationTime(time.Minute*2),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间:30天，
		rotatelogs.WithMaxAge(time.Hour*24*Log_SAVE_DAYS),
		// WithRotationCount设置文件清理前最多保存的个数。
		//rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		logger.Errorf("config local file system for logger error: %v", err)
	}

	logger.SetLevel(logLevel)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//&logrus.JSONFormatter{
	//	TimestampFormat:"2006-01-02 15:04:05",
	//}
	return lfsHook
}

func NewStdoutHook() logrus.Hook {

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: os.Stdout,
		logrus.InfoLevel:  os.Stdout,
		logrus.WarnLevel:  os.Stdout,
		logrus.ErrorLevel: os.Stdout,
		logrus.FatalLevel: os.Stdout,
		logrus.PanicLevel: os.Stdout,
	}, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return lfsHook
}

//// demo example from https://github.com/bingerambo/lfshook
//var Log *logrus.Logger
//
//func NewLogger() *logrus.Logger {
//	if Log != nil {
//		return Log
//	}
//
//	path := "/var/log/go.log"
//	writer, _ := rotatelogs.New(
//		path+".%Y%m%d%H%M",
//		rotatelogs.WithLinkName(path),
//		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
//		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
//	)
//
//	logrus.AddHook(lfshook.NewHook(
//		lfshook.WriterMap{
//			logrus.InfoLevel:  writer,
//			logrus.ErrorLevel: writer,
//		},
//		&logrus.JSONFormatter{},
//	))
//
//	pathMap := lfshook.PathMap{
//		logrus.InfoLevel:  "/var/log/info.log",
//		logrus.ErrorLevel: "/var/log/error.log",
//	}
//	Log = logrus.New()
//	Log.Hooks.Add(lfshook.NewHook(
//		pathMap,
//		&logrus.JSONFormatter{},
//	))
//
//	return Log
//}