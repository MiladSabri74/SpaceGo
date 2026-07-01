// Package weather is about forecast weather.
package weather

// CurrentCondition is presenting current weather.
var CurrentCondition string
// CurrentLocation is presenting current location of weather.
var CurrentLocation string

// Forecast is return the weather of current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
