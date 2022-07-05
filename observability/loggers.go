package observability

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func JSONLogger(lvl logger.Level) logger.FieldLogger {
	l := logrus.New()
	l.Level = lvl
	l.SetFormatter(&logrus.JSONFormatter{})
	l.SetOutput(os.Stdout)
	return logger.Logrus{FieldLogger: l}
}

func Log(app *buffalo.App, err error, lvl logrus.Level, message string) {
	if err != nil {
		switch lvl {
		case logger.DebugLevel:
			app.Logger.Debugf("%s: %v", message, err)
		case logger.InfoLevel:
			app.Logger.Infof("%s: %v", message, err)
		case logger.WarnLevel:
			app.Logger.Warnf("%s: %v", message, err)
		case logger.ErrorLevel:
			app.Logger.Errorf("%s: %v", message, err)
		default:
			app.Logger.Warnf("%s: %v", message, err)
		}
	}
}

func Panic(app *buffalo.App, err error) {
	if err != nil {
		app.Logger.Panic(err)
	}
}
