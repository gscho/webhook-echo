package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	level := os.Getenv("LOG_LEVEL")
  if level != "" {
  	switch strings.ToLower(level) {
  		case "info" : {
  			zerolog.SetGlobalLevel(zerolog.InfoLevel)
  		}
  		case "debug" : {
  			zerolog.SetGlobalLevel(zerolog.DebugLevel)
  		}
  		case "warn" : {
  			zerolog.SetGlobalLevel(zerolog.WarnLevel)
  		}
  		case "error" : {
  			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
  		}
  		case "default" :{
  			zerolog.SetGlobalLevel(zerolog.DebugLevel)
  		}
  	}
  }
	router := gin.Default()

	router.HEAD("/", func(c *gin.Context) {})
	router.POST("/webhook", func(c *gin.Context) {
		if jsonData, err := c.GetRawData(); err != nil {
			log.Error().Err(err).Msg("failed to get stream data")
			c.String(http.StatusBadRequest, "failed to get stream data")
		} else {
			if !json.Valid(jsonData) {
				log.Error().RawJSON("raw-json", jsonData).Msg("invalid json received")
				c.String(http.StatusBadRequest, "invalid json received")
			} else {
				log.Debug().RawJSON("raw-json", jsonData).Msg("success")
				c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
			}
		}
	})

	router.Run()
}
