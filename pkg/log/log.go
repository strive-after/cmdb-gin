package log

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	//file-rotatelogs 是一个日志轮转的库
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	//三个组件 一个日志打印  一个把日志写入文件 一个把文件作分割

	"gin-moudle/internal/config"
)

var (
	glog = logrus.New()
	//将字段的结构添加到日志条目。 它所做的只是调用WithWith来
	//每个`Field`。
	Log = glog.WithFields(logrus.Fields{
		"role": "cmdb",
	})
	normalWriter io.Writer
	hostName, _  = os.Hostname()
)

func InitLoger(c *config.Log) func() {

	formatter := &DefaultFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		HostName:        hostName,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			//findCaller 打印文件  行数 函数名
			return "", findCaller()
		},
	}
	// SetOutput设置记录器输出。
	glog.SetOutput(ioutil.Discard)

	glog.SetFormatter(
		//TextFormatter格式记录到文本中
		&logrus.TextFormatter{
			DisableColors:   true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05.000",
		},
	)

	normalWriter, err := rotatelogs.New(
		filepath.Join(c.Path, c.Filename+"_%Y-%m-%d.log"),
		//WithLinkName 这个是创建一个软连接 连接到当前日志
		rotatelogs.WithLinkName(filepath.Join(c.Path, c.Filename)),
		// WithMaxAge创建一个新的Option，该选项设置从文件系统清除日志文件之前的最长期限。 也就是保留历史日志日期
		rotatelogs.WithMaxAge(time.Duration(c.Maxage)*24*time.Hour),
		// WithRotationTime创建一个新的Option，以设置两次旋转之间的时间。 也就是切割日期
		rotatelogs.WithRotationTime(time.Duration(c.Rotation)*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			logrus.DebugLevel: normalWriter,
			logrus.InfoLevel:  normalWriter,
			logrus.WarnLevel:  normalWriter,
			logrus.ErrorLevel: normalWriter,
			logrus.FatalLevel: normalWriter,
			logrus.PanicLevel: normalWriter,
		},
		&logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05"},
	)

	lfHook.SetFormatter(formatter)
	switch c.Level {
	case "debug":
		glog.SetLevel(logrus.DebugLevel)
	case "info":
		glog.SetLevel(logrus.InfoLevel)
	case "warn":
		glog.SetLevel(logrus.WarnLevel)
	case "error":
		glog.SetLevel(logrus.ErrorLevel)
	case "fatal":
		glog.SetLevel(logrus.FatalLevel)
	case "panic":
		glog.SetLevel(logrus.PanicLevel)
	default:
		glog.SetLevel(logrus.InfoLevel)
	}

	glog.AddHook(lfHook)

	return Flush
}

func Flush() {
	if buffWriter, ok := normalWriter.(*bufio.Writer); ok {
		buffWriter.Flush()
	}
}

type DefaultFormatter struct {
	TimestampFormat  string
	HostName         string
	CallerPrettyfier func(f *runtime.Frame) (string, string)
}

func findCaller() string {
	var (
		funcName = ""
		file     = ""
		line     = 0
		pc       uintptr
	)

	// logrus + lfshook + log.go = 12
	for i := 10; i < 15; i++ {
		file, line, pc = getCaller(i)
		// fileter logrus + lfshook + log.go
		if strings.HasPrefix(file, "log/log.go") {
			continue
		}
		if strings.HasPrefix(file, "logrus") {
			continue
		}
		if strings.Contains(file, "lfshook.go") {
			continue
		}
		break
	}

	fullFnName := runtime.FuncForPC(pc)
	//fmt.Println(fullFnName)
	if fullFnName != nil {
		fnNameStr := fullFnName.Name()
		parts := strings.Split(fnNameStr, ".")
		funcName = parts[len(parts)-1]
	}
	//fmt.Printf("%s:%d:%s()\n", file, line, funcName)
	return fmt.Sprintf("%s:%d:%s()", file, line, funcName)
}

func getCaller(skip int) (string, int, uintptr) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0, pc
	}

	n := 0

	// get package name
	//fmt.Println(pc, file, line)
	//20411945 /Users/qushuaibo/Desktop/cmdb-gin/cmd/main.go 29
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	//20411945 cmd/main.go 29
	//fmt.Println(pc, file, line)
	return file, line, pc
}

func (f *DefaultFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// time field
	b.WriteString(entry.Time.Format(f.TimestampFormat))

	// hostname field
	b.WriteString("$$" + f.HostName)

	// level field
	b.WriteString("$$" + strings.ToUpper(entry.Level.String()))

	// component field
	b.WriteString("$$" + entry.Data["role"].(string))

	_, fileVal := f.CallerPrettyfier(entry.Caller)
	// file no
	b.WriteString("$$" + fileVal)

	// msg field
	b.WriteString("$$" + entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

func Errorf(format string, args ...interface{}) {
	Log.Errorf(format, args...)
}

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Warnf(format string, args ...interface{}) {
	Log.Warnf(format, args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Infof(format string, args ...interface{}) {
	Log.Infof(format, args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Printf(format string, args ...interface{}) {
	Log.Printf(format, args...)
}

func Print(args ...interface{}) {
	Log.Print(args...)
}

func Debugf(format string, args ...interface{}) {
	Log.Debugf(format, args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Panicf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(99)
}

func Panic(mesg, err error) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("mesg: %s, err: %s", mesg, err))
	os.Exit(99)
}

func PanicError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(99)
}
