// Package logger provides structured, rotating loggers shared across the app.
package logger

import (
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

const logDir = "var/logs"

// New returns a slog.Logger that writes JSON lines to logs/<name>.log,
// rotating the file by size/age and keeping a bounded number of backups.
func New(name string) *slog.Logger {
	if err := os.MkdirAll(logDir, 0o755); err != nil {
		slog.Error("failed to create log directory, falling back to stderr", "dir", logDir, "error", err)
		return slog.New(slog.NewJSONHandler(os.Stderr, nil))
	}

	rotator := &lumberjack.Logger{
		Filename:   filepath.Join(logDir, name+".log"),
		MaxSize:    50, // MB
		MaxBackups: 5,
		MaxAge:     30, // days
		Compress:   true,
	}

	return slog.New(slog.NewJSONHandler(rotator, nil))
}
