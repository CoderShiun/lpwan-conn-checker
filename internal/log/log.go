package log

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gitlab.com/MXCFoundation/cloud/conn_checker/internal/config"
	"path"
	"time"
)

var logPath = "."
var logName = "conn-checker"

func SetUp() error{
	logrus.SetLevel(logrus.Level(uint8(config.Conf.General.LogLevel)))
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := ConfigLocalFilesystemLogger(); err != nil {
		logrus.WithError(err).Error("Cannot hook lfshook")
		return err
	}

	return nil
}

func ConfigLocalFilesystemLogger() error{
	baseLogPaht := path.Join(logPath, logName)
	writer, err := rotatelogs.New(
		baseLogPaht+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPaht), // link to the logfile
		//rotatelogs.WithMaxAge(maxAge), // how long it's gonna save
		rotatelogs.WithRotationTime(time.Hour * 24), // when is gonna divide the file
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
		return err
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		//logrus.DebugLevel: writer, // set up different log for different err level
		//logrus.InfoLevel:  writer,
		//logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{DisableColors: true})

	logrus.AddHook(lfHook)

	return nil
}