package main

import (
	"os"
	"strings"

	"github.com/gscho/webhook-echo/internal/router"
	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	level := os.Getenv("LOG_LEVEL")
	if level != "" {
		switch strings.ToLower(level) {
		case "info":
			{
				zerolog.SetGlobalLevel(zerolog.InfoLevel)
			}
		case "warn":
			{
				zerolog.SetGlobalLevel(zerolog.WarnLevel)
			}
		case "error":
			{
				zerolog.SetGlobalLevel(zerolog.ErrorLevel)
			}
		}
	}
}

func main() {
	if err := router.Start(); err != nil {
		panic(err)
	}
}
