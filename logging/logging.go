package logging

import (
	"fmt"
	"github.com/semmisha/ClientAPI"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"io"
	"os"
	"path"
	"runtime"
)

func Logger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.ReportCaller = true
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             true,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          true,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier: func(frm *runtime.Frame) (function string, file string) {
			file = path.Base(frm.File)
			return fmt.Sprintf("%s", frm.Function), fmt.Sprintf("file:%v , line:%v", file, frm.Line)

		},
	})
	logger.SetOutput(io.Discard)

	tlgWriter := ClientAPI.NewConnection("@test111_111", "http://127.0.0.1:9999/PostMessage")
	filename := "/app/logging/logs.txt"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		logger.Error("Unable to Create/open log file:", filename)
	}

	logger.AddHook(&writer.Hook{
		Writer:    io.MultiWriter(file, os.Stdout),
		LogLevels: logrus.AllLevels,
	})
	logger.AddHook(&writer.Hook{
		Writer:    tlgWriter,
		LogLevels: []logrus.Level{logrus.FatalLevel, logrus.PanicLevel, logrus.ErrorLevel},
	})

	return logger
}
