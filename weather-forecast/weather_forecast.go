// Package weather provides the current weather conditions at various locations.
package weather

// CurrentCondition represents the current weather conditions at a specific location.
var CurrentCondition string

// CurrentLocation represents a city whose current weather conditions are of interest.
var CurrentLocation string

// Forecast returns a string describing the given city and its current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
