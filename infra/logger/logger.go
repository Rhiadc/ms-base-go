package logger

import (
	"log/slog"
	"os"
	"sync"

	"github.com/Rhiadc/ms-base-go/config"
)

var Log *Logger
var once sync.Once

type Logger struct {
	level    string
	devMode  bool
	encoding string
	Logger   *slog.Logger
}

func NewLogger(appConfig config.Config) *Logger {
	return &Logger{
		level:    appConfig.ENV,
		devMode:  appConfig.Log.DevMode,
		encoding: appConfig.Log.Encoding,
	}
}

func (l *Logger) InitLogger() {
	var opts *slog.HandlerOptions
	if l.devMode {
		opts = &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	l.Logger = logger
	once.Do(func() {
		Log = l
	})
}

func GetLogger() *Logger {
	return Log
}
