package logv

import "testing"

func TestNew(t *testing.T) {
	f := &Format{
		Filename:   "./a.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 3,
		Loglevel:   DebugLevel,
	}
	logger := New(f)
	for i := 0; i < 100000; i++ {
		logger.Warnln(i)
		logger.Traceln(i)
		logger.Debugln(i)
		logger.Errorln(i)
		logger.Infoln(i)
	}
}
