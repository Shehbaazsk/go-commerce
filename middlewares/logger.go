package middlewares

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
)

func CustomLogger() gin.HandlerFunc {
	// Create logs directory if not exists
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0755)
	}

	// Get today's date string
	dateStr := time.Now().Format("02-01-2006")

	// Rotating log file setup
	logFile := &lumberjack.Logger{
		Filename:   fmt.Sprintf("logs/log-%s.log", dateStr),
		MaxSize:    10,   // megabytes
		MaxBackups: 0,    // NO limit
		MaxAge:     30,   // days
		Compress:   true, // gzip old logs
	}

	// Combine file + stdout
	gin.DefaultWriter = io.MultiWriter(os.Stdout, logFile)

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf(`{"time":"%s","status":%d,"method":"%s","path":"%s","latency(In Sec)":"%f","ip":"%s"}`+"\n",
			param.TimeStamp.Format("02-01-2006 03:04:05 PM"),
			param.StatusCode,
			param.Method,
			param.Path,
			param.Latency.Seconds(),
			param.ClientIP,
		)
	})
}
