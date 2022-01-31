package booking

import (
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const timeLayout string = "1/2/2006 15:04:05"
	t, _ := time.Parse(timeLayout, date)

	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const timeLayout string = "January 2, 2006 15:04:05"
	t, _ := time.Parse(timeLayout, date)

	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	const timeLayout string = "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(timeLayout, date)

	return t.Hour() > 11 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	// Mon Jan 2 15:04:05 -0700 MST 2006
	var formattedDay string = Schedule(date).Format("Monday, January 2, 2006")
	var formattedTime string = Schedule(date).Format("15:04")

	return "You have an appointment on " + formattedDay + ", at " + formattedTime + "."
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
