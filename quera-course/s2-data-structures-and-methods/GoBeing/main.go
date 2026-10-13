package main

import "fmt"

func FormatTimeLimit(hour, minute, second int) string {
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func ProcessScreenTime(totalSeconds int) (hour int, minute int, second int) {
	hour = totalSeconds / (60 * 60)
	totalSeconds = totalSeconds - (hour * 60 * 60)
	minute = totalSeconds / 60
	second = totalSeconds - (minute * 60)
	return hour, minute, second
}
