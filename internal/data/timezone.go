package data

import (
	"os"
	"strings"
)

type TimeZoneService struct {
	// time zone directory might be different for different systems
	// timeZoneDir string
}

// Shameless copy from stack overflow https://stackoverflow.com/questions/40120056/get-a-list-of-valid-time-zones-in-go
func (tz *TimeZoneService) LoadTimeZones(path string) []string {
	var timeZones []string
	zoneDir := "/usr/share/zoneinfo/"
	files, _ := os.ReadDir(zoneDir + path)
	for _, f := range files {
		if f.Name() != strings.ToUpper(f.Name()[:1])+f.Name()[1:] {
			continue
		}
		if f.IsDir() {
			timeZones = append(timeZones, tz.LoadTimeZones(path+"/"+f.Name())...)
			tz.LoadTimeZones(path + "/" + f.Name())
		} else {
			timeZones = append(timeZones, (path + "/" + f.Name())[1:])
		}
	}
	return timeZones
}
