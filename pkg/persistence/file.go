package persistence

import (
	"fmt"
	"os"

	ics "github.com/arran4/golang-ical"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

// var PathTree defines where calendars will be stored
// This variable will be initialized in the main function
var PathTree string

// func WriteCalendar will persist a given calendar for a gesamtspielplan
// All structure (directory and files) will be generated.
// The function will return a string with the path of the calendar and error.
func WriteCalendar(gsp sport.Gesamtspielplan, cal *ics.Calendar) (string, error) {
	path := fmt.Sprintf("%s/%s/%s/",
		PathTree,
		gsp.Season,
		gsp.Championship.GetAbbreviation(),
	)
	filename := fmt.Sprintf("%s.ics", gsp.Group.String())

	os.MkdirAll(path, 0755)

	err := os.WriteFile(path+filename, []byte(cal.Serialize()), 0644)
	if err != nil {
		return path, err
	}

	return path, nil
}
