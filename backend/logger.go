// Logger nivelado sobre log/slog. Helpers printf-style para mantener los
// call-sites cortos (el handler text de slog ya añade time + level + msg).
//
// Nivel configurable por bandera -log-level o env SPECTRA_LOG_LEVEL.
// Valores: debug | info | warn | error.

package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

var logger *slog.Logger

func setupLogger(level string) {
	var lv slog.Level
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug":
		lv = slog.LevelDebug
	case "warn", "warning":
		lv = slog.LevelWarn
	case "error":
		lv = slog.LevelError
	default:
		lv = slog.LevelInfo
	}
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: lv}))
	slog.SetDefault(logger)
}

func logDebugf(format string, args ...any) {
	if logger.Enabled(context.Background(), slog.LevelDebug) {
		logger.Debug(fmt.Sprintf(format, args...))
	}
}

func logInfof(format string, args ...any) {
	if logger.Enabled(context.Background(), slog.LevelInfo) {
		logger.Info(fmt.Sprintf(format, args...))
	}
}

func logWarnf(format string, args ...any) {
	if logger.Enabled(context.Background(), slog.LevelWarn) {
		logger.Warn(fmt.Sprintf(format, args...))
	}
}

func logErrorf(format string, args ...any) {
	if logger.Enabled(context.Background(), slog.LevelError) {
		logger.Error(fmt.Sprintf(format, args...))
	}
}

func logFatalf(format string, args ...any) {
	logger.Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}
