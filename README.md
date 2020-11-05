# logrusdemo

## 1. Based on logrus, Add LFS hook to encapsulated to realize rotatation funtional log which can compress and delete old log entries
## 2. logrus version: v1.7.0 
## 3. more examples come from [logrus][logrus]
[logrus]: https://github.com/sirupsen/logrus

## 4. demo log print 
```
time="2020-11-05 16:41:49" level=info msg="hello world" size=10 zzz_line="cmd/main.go:18"
time="2020-11-05 16:41:49" level=warning msg="hello warn" podID=XXXUSXSSssaaf12334 zzz_line="cmd/main.go:19"
time="2020-11-05 16:41:49" level=error msg="Oh, error: &{Tom 111 12}" zobject="&{Tom 111 12}" zzz_line="cmd/main.go:21"
time="2020-11-05 16:41:49" level=info msg="Hi: {name:ids id:1 age:8}" ID-Object="&{Tom 111 12}" age=11 id=1 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:50" level=info msg="Hi: {name:ids id:2 age:8}" ID-Object="&{Tom 111 12}" age=11 id=2 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:51" level=info msg="Hi: {name:ids id:3 age:8}" ID-Object="&{Tom 111 12}" age=11 id=3 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:52" level=info msg="Hi: {name:ids id:4 age:8}" ID-Object="&{Tom 111 12}" age=11 id=4 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:53" level=info msg="Hi: {name:ids id:5 age:8}" ID-Object="&{Tom 111 12}" age=11 id=5 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:54" level=info msg="Hi: {name:ids id:6 age:8}" ID-Object="&{Tom 111 12}" age=11 id=6 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:55" level=info msg="Hi: {name:ids id:7 age:8}" ID-Object="&{Tom 111 12}" age=11 id=7 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:56" level=info msg="Hi: {name:ids id:8 age:8}" ID-Object="&{Tom 111 12}" age=11 id=8 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:57" level=info msg="Hi: {name:ids id:9 age:8}" ID-Object="&{Tom 111 12}" age=11 id=9 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:58" level=info msg="Hi: {name:ids id:10 age:8}" ID-Object="&{Tom 111 12}" age=11 id=10 zzz_line="cmd/main.go:23"
time="2020-11-05 16:41:59" level=info msg="Hi: {name:ids id:11 age:8}" ID-Object="&{Tom 111 12}" age=11 id=11 zzz_line="cmd/main.go:23"
time="2020-11-05 16:42:00" level=info msg="Hi: {name:ids id:12 age:8}" ID-Object="&{Tom 111 12}" age=11 id=12 zzz_line="cmd/main.go:23"
time="2020-11-05 16:42:01" level=info msg="Hi: {name:ids id:13 age:8}" ID-Object="&{Tom 111 12}" age=11 id=13 zzz_line="cmd/main.go:23"
time="2020-11-05 16:42:02" level=info msg="Hi: {name:ids id:14 age:8}" ID-Object="&{Tom 111 12}" age=11 id=14 zzz_line="cmd/main.go:23"

```