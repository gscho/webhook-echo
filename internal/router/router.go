package router

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	r.HEAD("/", head)
	r.POST("/webhook", handleWebhook)
}

func Start() error {
	return r.Run()
}

func head(c *gin.Context) {}

func handleWebhook(c *gin.Context) {
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
}
