package elon

import "fmt"

// Drive increases the car's distance driven by its speed and reduces its
// battery according to the battery drainage.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

// DisplayDistance returns the LED display for the distance driven.
func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the LED display for the remaining battery charge.
func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish returns true if the car can finish a race of the given distance.
func (c Car) CanFinish(trackDistance int) bool {
	return c.speed*(c.battery/c.batteryDrain) >= trackDistance
}
