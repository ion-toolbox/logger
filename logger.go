package logger

import (
	"fmt"
	joonix "github.com/joonix/log"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
)

type Logger struct {
	entry  *log.Entry
	logger *log.Logger
}

var instance *Logger

func New() *Logger {
	lgr := log.New()
	return &Logger{
		logger: lgr,
		entry:  lgr.WithFields(log.Fields{}),
	}
}

func Shared() *Logger {
	if instance == nil {
		instance = New()
		if strings.ToLower(os.Getenv("LOG_FORMAT")) == "kubernetes" {
			instance.logger.SetFormatter(joonix.NewFormatter())
		}
		if strings.ToLower(os.Getenv("LOG_FORMAT")) == "json" {
			instance.logger.SetFormatter(&log.JSONFormatter{})
		}
		switch strings.ToLower(os.Getenv("LOG_LEVEL")) {
		case "debug":
			instance.logger.SetLevel(log.DebugLevel)
		case "info":
			instance.logger.SetLevel(log.InfoLevel)
		case "warn":
			instance.logger.SetLevel(log.WarnLevel)
		case "err":
			instance.logger.SetLevel(log.ErrorLevel)
		}
	}
	return instance
}

func (l *Logger) GetWriter() io.Writer {
	return l.entry.Writer()
}

func (l *Logger) Debug(format string) {
	l.entry.Debug(format)
}

func (l *Logger) Debugf(format string, params ...interface{}) {
	l.entry.Debugf(format)
}

func (l *Logger) Info(format string) {
	l.entry.Info(format)
}

func (l *Logger) Infof(format string, params ...interface{}) {
	l.entry.Infof(format, params...)
}

func (l *Logger) Warning(format string) {
	l.entry.Warn(format)
}

func (l *Logger) Warningf(format string, params ...interface{}) {
	l.entry.Warnf(format, params...)
}

func (l *Logger) Error(format string) {
	l.entry.Error(format)
}

func (l *Logger) Errorf(format string, params ...interface{}) {
	l.entry.Errorf(format, params...)
}

func (l *Logger) Fatal(format string) {
	l.entry.Log(log.FatalLevel, format)
}

func (l *Logger) Fatalf(format string, params ...interface{}) {
	l.entry.Fatalf(format, params...)
}

func GetWriter() io.Writer {
	return Shared().entry.Writer()
}

func Debug(format string) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Debug(format)
}

func Debugf(format string, params ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Debugf(format, params...)
}

func Info(format string) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Info(format)
}

func Infof(format string, params ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Infof(format, params...)
}

func Warning(format string) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Warn(format)
}

func Warningf(format string, params ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Warnf(format, params...)
}

func Error(format string) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Error(format)
}

func Errorf(format string, params ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Errorf(format, params...)
}

func Fatal(format string) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Fatal(format)
}

func Fatalf(format string, params ...interface{}) {
	_, filename, line, _ := runtime.Caller(1)
	Shared().entry.WithFields(log.Fields{"at": fmt.Sprintf("%s:%d", filename, line)}).Fatalf(format, params...)
}
