package zaplogger

import "testing"

func TestNewLogger(t *testing.T) {
	zc, err := NewConfig(ZapConfig{
		Development:  true,
		Encoder:      "console",
		Level:        "debug",
		StackLevel:   "error",
		Sample:       false,
		TimeEncoding: "iso8601",
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
