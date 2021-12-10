package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuCal/pkg/calendar"
	"github.com/taskmedia/nuCal/pkg/persistence"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

// addRouterGesamtspielplan will add a REST endpoint to generate ICS / iCal files from a Gesamtspielplan
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

		// create calendars
		calGroup, calTeams := calendar.ConvertGesamtspielplanToGroupAndTeamCalendars(gsp)

		// persist calendars
		err := persistence.WriteCalendars(gsp, calGroup, calTeams)
		if err != nil {
			msg := "Could not write calendars to filesystem"
			log.WithFields(log.Fields{
				"gin.context": c,
			}).Warn(msg)
			c.String(http.StatusInternalServerError, msg)
			return
		}

		c.String(http.StatusCreated, "calandar created / updated")
	})
}
