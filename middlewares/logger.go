package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

// Custom Gin Logger with file rotation
func CustomGinLogger() gin.HandlerFunc {
	todayDate := time.Now().Format("02-01-2006")
	logFilename := fmt.Sprintf("logs/app_%s.log", todayDate)

	// Create a lumberjack logger for log file rotation
	logFile := &lumberjack.Logger{
		Filename:   logFilename, // Log file for Gin requests
		MaxSize:    10,          // Max size 10MB
		MaxBackups: 30,          // Keep 30 backups
		MaxAge:     1,           // Retain logs for 1 days
		Compress:   true,        // Compress log files
	}

	// Set the Gin log writer to our lumberjack log file
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	// Create a custom log format for Gin
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			// Custom log formatting
			return fmt.Sprintf(
				`{"ts":"%s","method":"%s","status_code":%d,"latency":"%.3fms","client_ip":"%s","url":"%s","error":"%s"}%s`,
				params.TimeStamp.Format("02/01/2006 15:04:05"),
				params.Method,
				params.StatusCode,
				float64(params.Latency.Microseconds())/1000.0,
				params.ClientIP,
				params.Request.URL,
				params.ErrorMessage,
				"\n",
			)
		},
		Output: logFile,
	})
}
