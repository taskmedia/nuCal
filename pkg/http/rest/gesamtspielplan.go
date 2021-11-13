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
		var matches sport.Matches

		if err := c.ShouldBindJSON(&matches); err != nil {
			msg := "The payload could not be binded to Matches object"
			log.WithField("gin.context", c).Warning(msg)
			c.String(http.StatusBadRequest, msg)
			return
		}

		fmt.Println(matches)

		c.String(http.StatusAccepted, "not yet implemented")
	})
}
