package logger

import "testing"

func TestLogger(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

func TestNewLogger(t *testing.T) {
	zc, err := NewConfig(&ZapConfig{
		Encoder:      EncoderConsole,
		Level:        LevelDebug,
		StackLevel:   LevelPanic,
		Sample:       false,
		TimeEncoding: TimeEncodingISO8601,
	})
	if err != nil {
		t.Fatal(err)
	}
	l := NewLogger(zc).Sugar()
	l.Debug("debug")
	l.Info("info")
	l.Warn("warn")
}

func TestSetStackEncoder(t *testing.T) {
	logger.Info("Test console log")
	SetEncoder(EncoderJSON)
	logger.Info("Test json log")
}
