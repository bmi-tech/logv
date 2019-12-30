package logv

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger is log type for package logv.
type Logger struct {
	// logrus.Logger is underline logger.
	*logrus.Logger

	// lumberjackLogger is used to rotate log file, lumberjackLogger is not set by default,
	// this will be set when SetOutput function is being called
	lumberjackLogger *lumberjack.Logger
}

// Format used to set up log file format
type Format struct {
	Filename string
	// MaxSize uses megabytes
	MaxSize int
	// MaxBackups set backup log file number
	MaxBackups int
	// MaxAge uses days, MaxAge set the lifetime of backup log file
	MaxAge int
	// Loglevel
	Loglevel Level
}

// Level is log level type
type Level uint32

// logging levels
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// NewDefault create new logger that output to stdout.
func NewDefault() *Logger {
	return &Logger{
		Logger: logrus.New(),
	}
}

// New create new logger that output to stdout.
// Default Value:
//         	Filename: ./logv.log
//			MaxSize: 50
//			MaxBackups: 3
//			MaxAge: 10
//			Loglevel: PanicLevel
func New(f *Format) *Logger {
	if len(f.Filename) == 0 {
		f.Filename = "./logv.log"
	}
	if f.MaxSize <= 0 {
		f.MaxSize = 50
	}
	if f.MaxBackups <= 0 {
		f.MaxBackups = 3
	}
	if f.MaxAge <= 0 {
		f.MaxAge = 10
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   f.Filename,
		MaxSize:    f.MaxSize,
		MaxBackups: f.MaxBackups,
		MaxAge:     f.MaxAge,
	}
	multiIO := io.MultiWriter(os.Stdout, lumberjackLogger)

	l := logrus.New()
	l.SetOutput(multiIO)
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	// Only log the warning severity or above.
	l.SetLevel(logrus.Level(f.Loglevel))

	return &Logger{
		Logger:           l,
		lumberjackLogger: lumberjackLogger,
	}
}
