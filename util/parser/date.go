package parser

import (
	"strings"
	"time"
)

func ParseTimeFromDB(rawDate string) (time.Time, error) {
	return time.Parse(time.RFC3339, strings.TrimSpace(rawDate))
}

func FormatDate(inputDate time.Time) string {
	return inputDate.Format("02 Jan 2006")
}

