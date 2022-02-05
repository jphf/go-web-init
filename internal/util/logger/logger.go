package logger

import (
	nested "github.com/antonfisher/nested-logrus-formatter"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Setup() {
	Logger = logrus.New()
	Logger.SetReportCaller(true)
	Logger.SetFormatter(&nested.Formatter{
		HideKeys: true,
	})
}
