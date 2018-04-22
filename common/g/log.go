package g

import (
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func InitLog() {
	conf := Conf()
	if !conf.Log.Enable {
		log.Info("log to std err")
		log.SetOutput(os.Stdout)
		log.SetLevel(log.DebugLevel)
		return
	}

	err := os.MkdirAll(conf.Log.Path, 0777)
	if err != nil {
		log.Fatalf("create directory %s failure\n", conf.Log.Path)
	}

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{ForceColors: true})

	logPath := conf.Log.Path
	maxAge := time.Duration(conf.Log.MaxAge)
	rotationTime := time.Duration(conf.Log.RotatTime)
	lfhook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: GetLogWriter(logPath, "debug", maxAge, rotationTime),
		log.InfoLevel:  GetLogWriter(logPath, "info", maxAge, rotationTime),
		log.WarnLevel:  GetLogWriter(logPath, "warn", maxAge, rotationTime),
		log.ErrorLevel: GetLogWriter(logPath, "error", maxAge, rotationTime),
		log.FatalLevel: GetLogWriter(logPath, "fatal", maxAge, rotationTime),
		log.PanicLevel: GetLogWriter(logPath, "panic", maxAge, rotationTime),
	}, &log.TextFormatter{ForceColors: true})
	log.AddHook(lfhook)
}

func GetLogWriter(logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) io.Writer {
	path := path.Join(getCurrPath(), logPath, logFileName)

	writer, err := rotatelogs.New(
		path+".%Y%m%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(maxAge*24*time.Hour),
		rotatelogs.WithRotationTime(rotationTime*24*time.Hour),
	)
	if err != nil {
		return nil
	}
	return writer
}

func getCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}
