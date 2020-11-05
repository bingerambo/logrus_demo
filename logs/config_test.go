package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)
func TestLogger_LogFn(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.WarnLevel)

	notCalled := 0
	logrus.InfoFn(func() []interface{} {
		notCalled++
		return []interface{}{
			"Hello",
		}
	})
	//assert.Equal(t, 0, notCalled)
	fmt.Println(notCalled)

	called := 0
	logrus.ErrorFn(func() []interface{} {
		called++
		return []interface{}{
			"Oopsi",
		}
	})
	//assert.Equal(t, 1, called)
	fmt.Println(called)
}

func TestExampleJSONFormatter_CallerPrettyfier(t *testing.T) {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Out = os.Stdout
	l.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return funcname, filename
		},
	}
	l.Info("example of custom format caller")
	// Output:
	// {"file":"example_custom_caller_test.go","func":"ExampleJSONFormatter_CallerPrettyfier","level":"info","msg":"example of custom format caller"}
}

