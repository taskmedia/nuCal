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

	cName := fmt.Sprintf("Handball: %s (%s %s %s)",
		gsp.AgeCategory.GetAbbreviation(),
		gsp.Championship,
		gsp.Class.GetAbbreviation(),
		gsp.Relay.GetAbbreviation(),
	)

	c.SetName(cName)
	c.SetXWRCalName(cName)

	cDesc := fmt.Sprintf("Handballkalender %s für die Spielklasse %s %s %s der Saison %s",
		gsp.AgeCategory.GetName(),
		gsp.Championship,
		gsp.Class.GetName(),
		gsp.Relay.GetName(),
		gsp.Season,
	)
	c.SetDescription(cDesc)

	c.SetMethod(ical.MethodPublish)

	gameDuration := time.Duration(float64(time.Hour) * 1.5)

	gspDesc := gsp.GetDescription()

	for _, m := range gsp.Matches {
		uid := fmt.Sprintf("%s-%s-%d-%d",
			gsp.Season,
			gsp.Championship,
			gsp.Group,
			m.Id)
		e := c.AddEvent(uid + "@nucal.task.media")

		eName := fmt.Sprintf("%s (%s %s %s): %s - %s",
			gsp.AgeCategory.GetAbbreviation(),
			gsp.Championship.GetAbbreviation(),
			gsp.Class.GetAbbreviation(),
			gsp.Relay.GetAbbreviation(),
			m.Team.Home,
			m.Team.Guest,
		)

		// add goals if available
		if m.Goal.Home != 0 {
			eName += fmt.Sprintf(" (%d:%d)", m.Goal.Home, m.Goal.Guest)
		}
		e.SetSummary(eName)
		e.SetModifiedAt(time.Now())
		e.SetStartAt(m.Date)
		e.SetEndAt(m.Date.Add(gameDuration))

		e.SetDescription(gspDesc + "\n" + m.GetDescription())

		if m.ReportId != 0 {
			url, _ := m.GetReportUrl()
			e.SetURL(url.String())
		}
	}

	return c
}
