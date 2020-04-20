package log

import (
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

type (

	// LoggerInterface is an interface that represent methods in log package
	LoggerInterface interface {
		Info(args ...interface{})
		Infof(format string, args ...interface{})
		Error(args ...interface{})
		Errorf(format string, args ...interface{})
	}

	// Logger return object of interface in  log package
	Logger struct {
		LoggerInterface
		prefix string
	}
)

var (
	logger   *logrus.Logger
	once     sync.Once
	instance *Logger
)

// Info is a logrus log message at level info on the standard logger
func (q *Logger) Info(args ...interface{}) {
	q.decorateLog().Info(args...)
}

// Infof is a logrus log message at level infof on the standard logger
func (q *Logger) Infof(format string, args ...interface{}) {
	q.decorateLog().Infof(format, args...)
}

// Error is a logrus log message at level error on the standard logger
func (q *Logger) Error(args ...interface{}) {
	q.decorateLog().Error(args...)
}

// Errorf is a logrus log message at level errorf on the standard logger
func (q *Logger) Errorf(format string, args ...interface{}) {
	q.decorateLog().Errorf(format, args...)
}

// NewLog is a factory that return  interface of log pakcage
func NewLog() *Logger {
	prefix := "experiment-server"
	once.Do(func() {
		logger = logrus.New()
		logger.Formatter = &prefixed.TextFormatter{
			FullTimestamp: true,
		}
		instance = &Logger{
			prefix: prefix,
		}
	})

	return instance
}

func (q *Logger) decorateLog() *logrus.Entry {
	var source string
	if pc, file, line, ok := runtime.Caller(2); ok {
		var funcName string
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
			if i := strings.LastIndex(funcName, "."); i != -1 {
				funcName = funcName[i+1:]
			}
		}

		source = fmt.Sprintf("%s:%v:%s()", path.Base(file), line, path.Base(funcName))
	}
	return logger.WithFields(logrus.Fields{
		"prefix": q.prefix,
		"source": source,
	})
}
