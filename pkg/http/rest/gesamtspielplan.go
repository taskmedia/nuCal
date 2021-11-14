package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

func addRouterGesamtspielplan(engine *gin.Engine) {
	engine.POST("/rest/v1/gesamtspielplan", func(c *gin.Context) {
		// validate request Content-Type
		contentType := c.Request.Header.Get("Content-Type")
		if contentType != "application/json" {
			msg := "Expected 'application/json' as content type"
			log.WithField("content-type", contentType).Warning(msg)
			c.String(http.StatusBadRequest, msg)
			return
		}

		var gsp sport.Gesamtspielplan

		if err := c.ShouldBindJSON(&gsp); err != nil {
			msg := "The payload could not be binded to Matches object"
			log.WithField("gin.context", c).Warning(msg)
			c.String(http.StatusBadRequest, msg)
			return
		}

		fmt.Println(gsp)

		c.String(http.StatusAccepted, "not yet implemented")
	})
}
