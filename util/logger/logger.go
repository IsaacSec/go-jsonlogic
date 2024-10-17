package logger

import (
	"fmt"
	"log/slog"
)

// ANSI color codes
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

var (
	logger *slog.Logger
)

func init() {

	logger = slog.New(slog.Default().Handler())
}

func Info(s string, v ...any) {
	prefix := Blue + "[INFO] " + Reset
	logger.Info(fmt.Sprintf(prefix+s, v...))
}

func Warn(s string, v ...any) {
	prefix := Yellow + "[WARN] " + Reset
	logger.Warn(fmt.Sprintf(prefix+s, v...))
}

func Error(s string, v ...any) {
	prefix := Red + "[ERROR] " + Reset
	logger.Error(fmt.Sprintf(prefix+s, v...))
}
