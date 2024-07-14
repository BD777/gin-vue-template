package logger

import (
	"context"
	"fmt"
	"gin-vue-template/pkg/utils/ctxmeta"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func init() {
	log.Printf("logger init\n")
	initLogrus()
}

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05.000")
	cwd, _ := os.Getwd()
	relativePath, _ := filepath.Rel(cwd, entry.Caller.File)
	msg := fmt.Sprintf("[%s %s %s %s:%d] %s\n",
		strings.ToUpper(entry.Level.String()),
		timestamp,
		path.Base(entry.Caller.Function),
		relativePath,
		entry.Caller.Line,
		entry.Message)
	return []byte(msg), nil
}

func initLogrus() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(new(MyFormatter))

	// exePath, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }

	// path := filepath.Join(filepath.Dir(exePath), "log/logrus.log")

	path := "log/logrus.log"
	log.Printf("log path: %v\n", path)

	// 创建一个新的 Writer，它会按日期分割日志文件，并自动删除7天前的文件
	writer, _ := rotatelogs.New(
		path+".%Y%m%d",
		rotatelogs.WithLinkName(path), // 生成软连接，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(24*7)*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour), // 日志切割时间间隔
	)
	log.Printf("current log file name: %v\n", writer.CurrentFileName())

	mw := io.MultiWriter(os.Stdout, writer)
	logrus.SetOutput(mw)
}

func f(ctx context.Context, format string) string {
	logid := ctxmeta.GetLogID(ctx)
	username := ctxmeta.GetUsername(ctx)
	return fmt.Sprintf("[logid:%s,user:%s] %s", logid, username, format)
}

func Info(ctx context.Context, format string, args ...interface{}) {
	logrus.Infof(f(ctx, format), args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	logrus.Errorf(f(ctx, format), args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	logrus.Warnf(f(ctx, format), args...)
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	logrus.Debugf(f(ctx, format), args...)
}
