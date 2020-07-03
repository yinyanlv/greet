package etc

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Log log.Logger

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.AddHook(newLfsHook(log.WarnLevel, 10000))
}

func newLfsHook(logLevel log.Level, maxRemainCount uint) log.Hook {
	logName := "log"
	writer, err := rotatelogs.New(
		path.Join("./" + AppConfig.Log.Dir, logName+".%Y%m%d%H"),

		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(path.Join("./" + AppConfig.Log.Dir, logName)),

		rotatelogs.WithRotationTime(time.Hour),

		// WithMaxAge和WithRotationCount二者只能设置一个，
		// WithMaxAge设置文件清理前的最长保存时间，
		// WithRotationCount设置文件清理前最多保存的个数。
		// rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCount),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	log.SetLevel(logLevel)

	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}
