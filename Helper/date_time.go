package Helper

import (
	"time"
)

func Date_Format(date string, format string) string {
	if date == "now" {
		t := time.Now()
		return t.Format(format)
	} else {
		t, _ := time.Parse(format, date)
		return t.Format(format)
	}
}
