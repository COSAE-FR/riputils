package logging

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

const defaultMaxSize = 10
const defaultMaxAge = 31
const defaultMaxBackups = 5

type Config struct {
	File           string `yaml:"file" json:"file" toml:"file"`
	Level          string `yaml:"level" json:"level" toml:"level"`
	FileMaxSize    int
	FileMaxAge     int
	FileMaxBackups int
	App            string `yaml:"-" json:"-" toml:"-"`
	Component      string `yaml:"-" json:"-" toml:"-"`
	Version        string `yaml:"-" json:"-" toml:"-"`
}

func SetupLog(config Config) (logger *log.Entry) {
	log.SetFormatter(&log.TextFormatter{
		ForceQuote:             true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		QuoteEmptyFields:       true,
	})
	if len(config.Level) == 0 {
		config.Level = "error"
	}
	logLevel, err := log.ParseLevel(config.Level)
	if err != nil {
		logLevel = log.ErrorLevel
	}
	log.SetLevel(logLevel)

	if len(config.File) > 0 {
		maxSize := defaultMaxSize
		maxAge := defaultMaxAge
		maxBackup := defaultMaxBackups
		if config.FileMaxSize > 0 {
			maxSize = config.FileMaxSize
		}
		if config.FileMaxAge > 0 {
			maxAge = config.FileMaxAge
		}
		if config.FileMaxBackups > 0 {
			maxBackup = config.FileMaxBackups
		}

		fileLog := &lumberjack.Logger{
			Filename:   config.File,
			MaxSize:    maxSize,
			MaxAge:     maxAge,
			MaxBackups: maxBackup,
			Compress:   true,
		}
		log.SetOutput(fileLog)
	} else {
		log.SetOutput(os.Stderr)
	}
	app := "rip"
	component := "logger_setup"
	if len(config.App) > 0 {
		app = config.App
	}
	if len(config.Component) > 0 {
		component = config.Component
	}
	logger = log.WithFields(log.Fields{
		"app":       app,
		"component": component,
	})
	if len(config.Version) > 0 {
		logger = logger.WithField("version", config.Version)
	}
	return
}
