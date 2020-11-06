package logs

import (
	filename "github.com/keepeye/logrus-filename"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	// linux
	LOG_FILE_PATHNAME = "/var/log/aistation/kube-batch/kubebatch"
	// windows
	//LOG_FILE_PATHNAME = "D:/GO_projects/src/github.com/bingerambo/logrus_demo/logs/kubebatch"
)

// global var
var (
	ALogger  *(logrus.Logger)
	nullfile *(os.File)
)

func StartUp() {
	ALogger = logrus.New()

	// desperate: 使用指定文件路径
	//ALogfile, err := os.OpenFile(LOG_FILE_PATHNAME, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	//if err == nil {
	//	ALogger.Out = ALogfile
	//} else {
	//	ALogger.Error("Failed to log to file, using default stderr")
	//}

	// make default output null, for linux and windows
	var err error
	nullfile, err = os.OpenFile(os.DevNull, os.O_WRONLY, 0666)
	if err != nil {
		ALogger.Fatalf("%v", err)
	}

	//defer outputfile.Close()
	ALogger.SetOutput(nullfile)

	filenameHook := filename.NewHook()
	filenameHook.Field = "zzz_line"

	// 设置filenamehook
	ALogger.AddHook(filenameHook)

	lfHook := NewLfsHook(ALogger, LOG_FILE_PATHNAME, logrus.InfoLevel, 60)

	// 设置lfshook
	ALogger.AddHook(lfHook)

	//设置stdoutHook
	stdoutHook := NewStdoutHook()
	ALogger.AddHook(stdoutHook)
}

func Stop() {
	if nullfile != nil {
		nullfile.Close()
	}
}
