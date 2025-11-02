// Package weather provides weather forecast.
package weather 

var (
    // CurrentCondition represents current weather condition.
	CurrentCondition string
    // CurrentLocation represents current location.
	CurrentLocation  string
)

// Forecast returns a string based on a city and its weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
