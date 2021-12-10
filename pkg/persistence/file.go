package persistence

import (
	"fmt"
	"os"
	"regexp"

	ics "github.com/arran4/golang-ical"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

// var PathTree defines where calendars will be stored
// This variable will be initialized in the main function
var PathTree string

// func WriteCalendar will persist a given calendar for a gesamtspielplan
// All structure (directory and files) will be generated.
// The function will return a string with the path of the calendar and error.
func WriteCalendars(gsp sport.Gesamtspielplan, groupCal *ics.Calendar, teamCal map[string]*ics.Calendar) error {
	dir := fmt.Sprintf("%s/%s/%s/%s/",
		PathTree,
		gsp.Season,
		gsp.Championship.GetAbbreviation(),
		gsp.Group.String(),
	)

	os.MkdirAll(dir, 0755)

	// persist group calendar
	absolutePath := dir + "group.ics"
	err := writeCalendar(absolutePath, groupCal)
	if err != nil {
		return err
	}

	// persist team calendars
	unify, _ := regexp.Compile("[^a-zA-Z0-9]+")
	for team, tc := range teamCal {
		filename := unify.ReplaceAllString(team, "")
		filename = fmt.Sprintf("%s.ics", filename)

		absolutePath := dir + filename
		err := writeCalendar(absolutePath, tc)
		if err != nil {
			return err
		}
	}
	return nil
}

// func writeCalendar wraps the file writing of a single calendar
func writeCalendar(absolutePath string, cal *ics.Calendar) error {
	return os.WriteFile(absolutePath, []byte(cal.Serialize()), 0644)
}
