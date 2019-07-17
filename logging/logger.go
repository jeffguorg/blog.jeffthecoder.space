package logging

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// SetupFormat setup time format of the logger
func SetupFormat() {
	formater := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15-04-05.999 MST",
	}
	logger.SetFormatter(formater)
	logger.Info("Logger has been setup")
}

// SetupLevel setup log level of the logger
func SetupLevel(v string) error {
	var (
		lvl logrus.Level
		err error
	)

	if lvl, err = logrus.ParseLevel(v); err != nil {
		return err
	}
	logger.SetLevel(lvl)

	return nil
}

func Fatal(args ...interface{}) {
	logger.Fatalln(args...)
}

func Error(args ...interface{}) {
	logger.Errorln(args...)
}
func Warn(args ...interface{}) {
	logger.Warnln(args...)
}
func Info(args ...interface{}) {
	logger.Infoln(args...)
}
func Debug(args ...interface{}) {
	logger.Debugln(args...)
}
