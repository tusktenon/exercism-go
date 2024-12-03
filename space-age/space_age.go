package space

const (
	mercuryYear = 0.2408467 * 31557600
	venusYear   = 0.61519726 * 31557600
	earthYear   = 31557600
	marsYear    = 1.8808158 * 31557600
	jupiterYear = 11.862615 * 31557600
	saturnYear  = 29.447498 * 31557600
	uranusYear  = 84.016846 * 31557600
	neptuneYear = 164.79132 * 31557600
)

type Planet string

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return seconds / mercuryYear
	case "Venus":
		return seconds / venusYear
	case "Earth":
		return seconds / earthYear
	case "Mars":
		return seconds / marsYear
	case "Jupiter":
		return seconds / jupiterYear
	case "Saturn":
		return seconds / saturnYear
	case "Uranus":
		return seconds / uranusYear
	case "Neptune":
		return seconds / neptuneYear
	default:
		return -1
	}
}
