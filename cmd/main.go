package main

import (
	"github.com/bingerambo/logrus_demo/logs"
	"github.com/sirupsen/logrus"
	"time"
)

type Student struct {
	name string
	id   int
	age  int
}

// more examples from https://github.com/sirupsen/logrus
func main() {
	logs.StartUp()
	logs.ALogger.WithField("size", 10).Info("hello world")
	logs.ALogger.WithField("podID", "XXXUSXSSssaaf12334").Warn("hello warn")
	a := &Student{"Tom", 111, 12}
	logs.ALogger.WithField("zobject", a).Errorf("Oh, error: %v", a)
	for num := 1; num < 50; num++ {
		logs.ALogger.WithFields(logrus.Fields{"ID-Object": a, "age": 11, "id": num}).Infof("Hi: %+v", Student{"ids", num, 8})
		time.Sleep(1 * time.Second)
	}
	logs.ALogger.WithFields(logrus.Fields{"zobject": a, "age": 11, "id": 211}).Errorf("Oh, error: %+v", Student{"Kate", 222, 10})
	logs.ALogger.WithField("size", 1000000).Info("log end ... ")
	logs.Stop()
}
