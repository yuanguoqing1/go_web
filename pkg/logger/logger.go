package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// 设置日志格式为 JSON 格式
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 创建 logs 目录
	if err := os.MkdirAll("logs", 0755); err != nil {
		fmt.Printf("create logs directory failed: %v\n", err)
		return
	}

	// 设置日志文件
	filename := path.Join("logs", fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open log file failed: %v\n", err)
		return
	}

	// 同时输出到文件和控制台
	Log.SetOutput(os.Stdout)
	Log.AddHook(&FileHook{
		Writer: file,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	})

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)
}

// FileHook 自定义 Hook，用于写入文件
type FileHook struct {
	Writer    *os.File
	LogLevels []logrus.Level
}

func (hook *FileHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *FileHook) Levels() []logrus.Level {
	return hook.LogLevels
}
