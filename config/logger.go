package config

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetLogger returns a logger instance with file rotation setup
func GetLogger() *zap.Logger {
	// Custom timestamp format for logs
	timeEncoder := zapcore.TimeEncoderOfLayout("02/01/2006 15:04:05")

	// Create a custom log format
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "ts",
		EncodeTime:   timeEncoder,
		EncodeLevel:  zapcore.CapitalLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// Set up log file rotation using lumberjack
	logFile := &lumberjack.Logger{
		Filename:   "logs/app.log", // Log file name
		MaxSize:    10,             // MB
		MaxBackups: 3,              // Number of backup files
		MaxAge:     7,              // Days to keep old log files
		Compress:   true,           // Compress old log files
	}

	// Create a core that writes logs to both file and console
	cores := []zapcore.Core{
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),                                             // JSON format
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logFile)), // Write to both stdout and log file
			zap.InfoLevel,
		),
	}

	// Create the logger
	logger := zap.New(zapcore.NewTee(cores...))
	return logger
}
