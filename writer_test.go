package logger

import (
	"fmt"
	"testing"
)

func TestNewLogWriter(t *testing.T) {
	w := NewLogWriter()

	_, err := fmt.Fprint(w, "Test log writer")
	if err != nil {
		t.Fatal(err)
	}
}
