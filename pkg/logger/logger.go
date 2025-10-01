package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var envLogLevel = map[string]zerolog.Level{
	"development": zerolog.DebugLevel,
	"staging":     zerolog.DebugLevel,
	"production":  zerolog.InfoLevel,
}

func Initialize(env string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	level, ok := envLogLevel[env]
	if !ok {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	log.Info().Msg("logger initialized")
}

func logBase(event *zerolog.Event, c *gin.Context, status int, latency float64) *zerolog.Event {
	return event.
		Str("method", c.Request.Method).
		Str("path", c.Request.URL.Path).
		Int("status", status).
		Str("latency_ms", fmt.Sprintf("%.3f", latency))
}

func ZerologMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		latency := float64(duration.Microseconds()) / 1000.0

		status := c.Writer.Status()
		var event *zerolog.Event
		switch {
		case status >= 500:
			event = log.Error()
		case status >= 400:
			event = log.Warn()
		default:
			event = log.Info()
		}

		base := logBase(event, c, status, latency)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				base.Err(err).Msg("error occurred")
			}
			return
		}

		base.Msg("request handled")
	}
}
