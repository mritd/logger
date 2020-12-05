package logger

import "testing"

func TestNewConfig(t *testing.T) {
	_, err := NewConfig(ZapConfig{
		Development:  true,
		Encoder:      "json",
		Level:        "debug",
		StackLevel:   "info",
		Sample:       false,
		TimeEncoding: "iso8601",
	})
	if err != nil {
		t.Fatal(err)
	}
}
