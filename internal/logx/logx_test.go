package logx_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/logx"
)

func TestLogger_Default(t *testing.T) {
	t.Parallel()

	l := logx.Default()
	if l == nil {
		t.Fatal("default logger is nil")
	}
}
