package raindrops

import "strconv"

func Convert(number int) (sound string) {
	if number%3 == 0 {
		sound = "Pling"
	}
	if number%5 == 0 {
		sound += "Plang"
	}
	if number%7 == 0 {
		sound += "Plong"
	}
	if sound == "" {
		sound = strconv.Itoa(number)
	}
	return
}
