package logger

import "testing"

func TestLogger(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

func TestNewLogger(t *testing.T) {
	zc, err := NewConfig(ZapConfig{
		Development:  true,
		Encoder:      EncoderConsole,
		Level:        LevelDebug,
		StackLevel:   LevelError,
		Sample:       false,
		TimeEncoding: TimeEncoderISO8601,
	})
	if err != nil {
		t.Fatal(err)
	}
	logger := NewLogger(zc)
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
