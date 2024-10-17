package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
)

type CustomHandler struct {
	out io.Writer
}

func (h *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	timeStr := r.Time.Format("2006-01-02 15:04:05")

	var levelColor string
	switch r.Level {
	case slog.LevelInfo:
		levelColor = "\033[34m" // Blue
	case slog.LevelWarn:
		levelColor = "\033[33m" // Yellow
	case slog.LevelError:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[0m" // Default color
	}

	timeColor := "\033[90m" // Grey
	resetColor := "\033[0m"

	fmt.Fprintf(h.out, "%s%s%s [%s%s%s] %s",
		timeColor, timeStr, resetColor,
		levelColor, level, resetColor,
		r.Message)

	// Handle additional attributes
	if r.NumAttrs() > 0 {
		fmt.Fprint(h.out, " {")
		r.Attrs(func(a slog.Attr) bool {
			fmt.Fprintf(h.out, "%s: %v, ", a.Key, a.Value)
			return true
		})
		fmt.Fprint(h.out, "}")
	}

	fmt.Fprintln(h.out) // Add a newline at the end
	return nil
}

func (h *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h // For simplicity, we're not handling WithAttrs
}

func (h *CustomHandler) WithGroup(name string) slog.Handler {
	return h // For simplicity, we're not handling WithGroup
}

func (h *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true // Enable all levels
}

func NewCustomLogger() *slog.Logger {
	handler := &CustomHandler{
		out: os.Stdout,
	}
	return slog.New(handler)
}

var (
	logger *slog.Logger
)

func init() {
	logger = NewCustomLogger()

	slog.SetDefault(logger)
}

func Info(s string, v ...any) {
	logger.Info(fmt.Sprintf(s, v...))
}

func Warn(s string, v ...any) {
	logger.Warn(fmt.Sprintf(s, v...))
}

func Error(s string, v ...any) {
	logger.Error(fmt.Sprintf(s, v...))
}
