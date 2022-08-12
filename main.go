package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	router := gin.Default()

	router.HEAD("/", func(c *gin.Context) {})
	router.POST("/webhook", func(c *gin.Context) {
		if jsonData, err := c.GetRawData(); err != nil {
			log.Print(err)
		} else {
			log.Debug().Msg(string(jsonData))
			c.Data(http.StatusOK, gin.MIMEJSON, jsonData)
		}
	})

	router.Run()
}
