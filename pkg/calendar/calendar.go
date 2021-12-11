package calendar

import (
	"fmt"
	"regexp"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

// unify will not allow special characters for string
// this can be used to remove illegal characters from teamname
var unify, _ = regexp.Compile("[^a-zA-Z0-9]+")

// func ConvertGesamtspielplanToGroupAndTeamCalendars will return calendars of the whole group and each team
// A given Gesamtspielplan will be converted to the groups calender which includes every game.
// Also (second return value) a calendar for each team will be generated.
func ConvertGesamtspielplanToGroupAndTeamCalendars(gsp sport.Gesamtspielplan) (*ics.Calendar, map[string]*ics.Calendar) {
	// declare calendar for group and each team
	calGroup := ics.NewCalendar()
	calTeams := make(map[string]*ics.Calendar)

	// calendars
	configureGesamptspielplanCalendarGroup(gsp, calGroup)
	dt := gsp.GetDistinctTeams()
	for _, team := range dt {
		calTeams[team] = ics.NewCalendar()
		configureGesamptspielplanCalendarTeam(gsp, calTeams[team], team)
	}

	gspDesc := gsp.GetDescription()

	for _, m := range gsp.Matches {
		// skip adding match if no date is not available
		if m.Date.IsZero() {
			continue
		}

		// create event
		e := createEventFromMatch(m, gsp, gspDesc)

		// add event to calendars (group, hometeam, guestteam)
		calGroup.Components = append(calGroup.Components, e)
		calTeams[m.Team.Home].Components = append(calTeams[m.Team.Home].Components, e)
		calTeams[m.Team.Guest].Components = append(calTeams[m.Team.Guest].Components, e)
	}

	return calGroup, calTeams
}

// func createEventFromMatch will create a standardized VEvent to
func createEventFromMatch(m sport.Match, gsp sport.Gesamtspielplan, gspDesc string) *ics.VEvent {
	// events prerequisites
	matchDuration := time.Duration(float64(time.Hour) * 1.5)

	uuid := fmt.Sprintf("%s-%s-%d-%d",
		gsp.Season,
		gsp.Championship.GetAbbreviation(),
		gsp.Group,
		m.Id)

	// create event with uuid
	e := ics.VEvent{
		ics.ComponentBase{
			Properties: []ics.IANAProperty{
				{ics.BaseProperty{IANAToken: ics.ToText(string(ics.ComponentPropertyUniqueId)), Value: uuid}},
			},
		},
	}

	summary := fmt.Sprintf("%s (%s %s %s): %s - %s",
		gsp.AgeCategory.GetAbbreviation(),
		gsp.Championship.GetAbbreviation(),
		gsp.Class.GetAbbreviation(),
		gsp.Relay.GetAbbreviation(),
		m.Team.Home,
		m.Team.Guest,
	)
	// add goals to summary if available
	if m.Goal.Home != 0 {
		summary += fmt.Sprintf(" (%d:%d)", m.Goal.Home, m.Goal.Guest)
	}
	e.SetSummary(summary)

	e.SetStartAt(m.Date)
	e.SetEndAt(m.Date.Add(matchDuration))

	e.SetDescription(gspDesc + "\n" + m.GetDescription())

	if m.ReportId != 0 {
		url, _ := m.GetReportUrl()
		e.SetURL(url.String())
	}

	e.SetModifiedAt(time.Now())

	return &e
}

// func configureGesamptspielplanCalendarGroup is a wrapper for the configuration for group calendars
func configureGesamptspielplanCalendarGroup(gsp sport.Gesamtspielplan, c *ics.Calendar) {
	configureGesamptspielplanCalendar(gsp, c, "")
}

// func configureGesamptspielplanCalendarGroup is a wrapper for the configuration for team calendars
func configureGesamptspielplanCalendarTeam(gsp sport.Gesamtspielplan, c *ics.Calendar, team string) {
	configureGesamptspielplanCalendar(gsp, c, team)
}

// func configureGesamptspielplanCalendar configures a calendar to Gesamtspielplan specifications
func configureGesamptspielplanCalendar(gsp sport.Gesamtspielplan, c *ics.Calendar, suffix string) {
	// calendar id
	prodid := fmt.Sprintf("%s-%s-%d",
		gsp.Season,
		gsp.Championship.GetAbbreviation(),
		gsp.Group,
	)

	// calendar name
	name := fmt.Sprintf("Handball: %s (%s %s %s)",
		gsp.AgeCategory.GetAbbreviation(),
		gsp.Championship.GetAbbreviation(),
		gsp.Class.GetAbbreviation(),
		gsp.Relay.GetAbbreviation(),
	)

	// calendar description
	desc := fmt.Sprintf("Handballkalender %s für die Spielklasse %s %s %s der Saison %s",
		gsp.AgeCategory.GetName(),
		gsp.Championship.GetName(),
		gsp.Class.GetName(),
		gsp.Relay.GetName(),
		gsp.Season,
	)

	if suffix != "" {
		suffixUnified := unify.ReplaceAllString(suffix, "")
		prodid += "-" + suffixUnified
		name += ": " + suffix
		desc += " für das Team " + suffix
	}

	// generate fullprodid
	prodid = fmt.Sprintf("-//nucal.task.media//%s//DE", prodid)

	// set values to calendar
	c.SetProductId(prodid)
	c.SetXWRCalID(prodid)

	c.SetName(name)
	c.SetXWRCalName(name)

	c.SetDescription(desc)

	c.SetMethod(ics.MethodPublish)
}
