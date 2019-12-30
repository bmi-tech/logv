package logv

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *Logger

func init() {
	logger = NewDefault()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
}

// SetLogger set default global logger config
func SetLogger(f *Format) {
	var filename string
	var maxSize, maxBackups, maxAge int
	var l Level

	if f.Filename == "" {
		filename = defaultFilename
	} else {
		filename = f.Filename
	}

	if f.MaxSize == 0 {
		maxSize = defaultMaxSize
	} else {
		maxSize = f.MaxSize
	}

	if f.MaxBackups == 0 {
		maxBackups = defaultMaxBackups
	} else {
		maxBackups = f.MaxBackups
	}

	if f.MaxAge == 0 {
		maxAge = defaultMaxAge
	} else {
		maxAge = f.MaxAge
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	multiIO := io.MultiWriter(os.Stdout, lumberjackLogger)
	logger.SetOutput(multiIO)
	logger.lumberjackLogger = lumberjackLogger
	logger.filename = filename

	logger.SetLevel(transLevelToLogrusLevel(l))
}

// SetOutputFile set default logger output file, if rotate is configed, reuse configuration,
// otherwise config use default value,
// 		MaxSize: 50 mega bytes
//		MaxBackups: 10
// 		MaxAge: 10 days
func SetOutputFile(filename string) {
	if logger.lumberjackLogger == nil {
		lumberjackLogger := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    defaultMaxSize,
			MaxBackups: defaultMaxBackups,
			MaxAge:     defaultMaxAge,
		}
		multiIO := io.MultiWriter(os.Stdout, lumberjackLogger)
		logger.SetOutput(multiIO)
		logger.lumberjackLogger = lumberjackLogger
		logger.filename = filename
	}
}

// SetRotate set log rotate info, if filename is configed, reuse configuration,
// otherwise config use default value,
// 		Filename: ./logv.log
func SetRotate(maxSize, maxBackups, maxAge int) {
	var lumberjackLogger *lumberjack.Logger
	if logger.filename == "" {
		lumberjackLogger = &lumberjack.Logger{
			Filename:   defaultFilename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		}
		logger.filename = defaultFilename
		multiIO := io.MultiWriter(os.Stdout, lumberjackLogger)
		logger.SetOutput(multiIO)
		logger.lumberjackLogger = lumberjackLogger
	} else {
		lumberjackLogger = &lumberjack.Logger{
			Filename:   logger.filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		}
		multiIO := io.MultiWriter(os.Stdout, lumberjackLogger)
		logger.SetOutput(multiIO)
		logger.lumberjackLogger = lumberjackLogger
	}
}

// SetLevel change log level
func SetLevel(l Level) {
	logger.SetLevel(transLevelToLogrusLevel(l))
}

// SetOutput sets the logger output.
// Be sure you want set your output, if you just want set logger
// output to a file, use SetOutputFile
func SetOutput(output io.Writer) {
	logger.SetOutput(output)
}

// Tracef see logrus Tracef
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

// Debugf see logrus Debugf
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof see logrus Infof
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf see logrus Warnf
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Errorf see logrus Errorf
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatalf will exit(1), see logrus Fatalf
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panicf will panic, see logrus Panicf
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}
