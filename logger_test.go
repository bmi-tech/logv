package logv

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	Tracef("trace %s", "trace")
	Debugf("Debugf %s", "Debugf")
	Infof("Infof %s", "Infof")
	Warnf("Warnf %s", "Warnf")
	Errorf("Errorf %s", "Errorf")
}

func TestSetOutputFile(t *testing.T) {
	SetOutputFile("./a.log")
	Tracef("trace %s", "trace")
}

func TestSetRotate(t *testing.T) {
	SetRotate(1, 2, 1)
	for i := 0; i < 100000; i++ {
		Debugf("%d", i)
	}
}

func TestSetOutput(t *testing.T) {
	f, err := os.Create("./b.log")
	if err != nil {
		t.Errorf("create b.log failed")
	}
	defer f.Close()
	SetOutput(f)
	Tracef("trace %s", "trace")
}

// func TestFatalf(t *testing.T) {
// 	Fatalf("Fatalf %s", "Fatalf")
// }

// func TestPanicf(t *testing.T) {
// 	Panicf("Panicf %s", "Panicf")
// }

func TestSetLogger(t *testing.T) {
	SetLogger(&Format{
		Filename:   "./b.log",
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 4,
		Loglevel:   InfoLevel,
	})
	Tracef("trace %s", "trace")
	Debugf("Debugf %s", "Debugf")
	Infof("Infof %s", "Infof")
	Warnf("Warnf %s", "Warnf")
	Errorf("Errorf %s", "Errorf")
}
