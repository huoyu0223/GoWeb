package logm

import (
	"GoWeb/src/conf"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func NewLogRus() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.Level(conf.LogCfg.Level))
	if conf.LogCfg.FileName != "" {
		filenameex := "logs/" + conf.LogCfg.FileName
		logger := &lumberjack.Logger{
			Filename:   filenameex,
			MaxSize:    conf.LogCfg.MaxFileSize, // megabytes
			MaxBackups: conf.LogCfg.MaxBackupIndex,
			MaxAge:     conf.LogCfg.MaxAge, // days
			Compress:   conf.LogCfg.Compress,
		}
		logrus.SetOutput(logger)
	}
}
