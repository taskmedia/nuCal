package calendar

import (
	"fmt"
	"time"

	ical "github.com/arran4/golang-ical"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

// ConvertGesamtspielplanToCalendar will convert a Gesamtspielplan struct into a calendar
func ConvertGesamtspielplanToCalendar(gsp sport.Gesamtspielplan) *ical.Calendar {
	c := ical.NewCalendar()

	cUid := fmt.Sprintf("-//nucal.task.media//%s-%s-%d//DE",
		gsp.Season,
		gsp.Championship,
		gsp.Group)
	c.SetProductId(cUid)
	c.SetXWRCalID(cUid)

	cName := fmt.Sprintf("%s (%d)",
		gsp.Championship,
		gsp.Group)
	c.SetName(cName)
	c.SetXWRCalName(cName)

	c.SetDescription("Handballkalender")
	c.SetRefreshInterval("Handballkalender")

	c.SetMethod(ical.MethodPublish)

	gameDuration := time.Duration(float64(time.Hour) * 1.5)

	for _, m := range gsp.Matches {
		uid := fmt.Sprintf("%s-%s-%d-%d",
			gsp.Season,
			gsp.Championship,
			gsp.Group,
			m.Id)
		e := c.AddEvent(uid + "@nucal.task.media")

		e.SetSummary(m.Team.Home + " - " + m.Team.Guest)
		e.SetModifiedAt(time.Now())
		e.SetStartAt(m.Date)
		e.SetEndAt(m.Date.Add(gameDuration))
	}

	return c
}
