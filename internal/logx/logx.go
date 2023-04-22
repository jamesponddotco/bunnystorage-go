// Package logx is a simple logger implementing the bunnystorage.Logger
// interface.
package logx

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

// Default represents the default logger used by our package.
type Logger struct {
	// logger is the underlying logger.
	logger *log.Logger

	// builder is used to build the log message.
	builder strings.Builder

	// mu is used to protect access to the builder.
	mu sync.Mutex
}

// Default returns the default logger.
func Default() *Logger {
	return &Logger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
}

// Printf calls the underlying logger's Printf method.
func (l *Logger) Printf(format string, v ...any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.builder.Reset()

	fmt.Fprintf(&l.builder, format, v...)

	l.logger.Print(l.builder.String())
}
