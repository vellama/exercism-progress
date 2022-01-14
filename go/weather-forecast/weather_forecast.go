// Package weather provides some tools to tell us more about the weather.
package weather

// CurrentCondition tells if it is cloudy, sunny, rainy, etc ...
var CurrentCondition string

// CurrentLocation can be the city, street, country, etc .....
var CurrentLocation string

// Forecast return current weather forecast for current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
