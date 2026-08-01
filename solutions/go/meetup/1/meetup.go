package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule string

const (
	First  WeekSchedule = "first"
	Second WeekSchedule = "second"
	Third  WeekSchedule = "third"
	Fourth WeekSchedule = "fourth"
	Last   WeekSchedule = "last"
	Teenth WeekSchedule = "teenth"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	// Get the number of days in the given month
	daysInMonth := daysInMonth(month, year)

	var matchingDays []int
	for day := 1; day <= daysInMonth; day++ {
		t := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
		if t.Weekday() == wDay {
			matchingDays = append(matchingDays, day)
		}
	}

	switch wSched {
	case First:
		return matchingDays[0]
	case Second:
		return matchingDays[1]
	case Third:
		return matchingDays[2]
	case Fourth:
		return matchingDays[3]
	case Last:
		return matchingDays[len(matchingDays)-1]
	case Teenth:
		for _, day := range matchingDays {
			if day >= 13 && day <= 19 {
				return day
			}
		}
	}

	return 0
}

func daysInMonth(month time.Month, year int) int {
	nextMonth := month + 1
	nextYear := year
	if nextMonth > time.December {
		nextMonth = time.January
		nextYear++
	}
	lastDay := time.Date(nextYear, nextMonth, 0, 0, 0, 0, 0, time.UTC)
	return lastDay.Day()
}
