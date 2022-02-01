package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	if c.batteryDrain > c.battery {
		return
	}

	c.distance += c.speed
	c.battery -= c.batteryDrain
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	var nbDrives int = 0
	var remainingDistance int = trackDistance - c.distance

	if remainingDistance <= 0 {
		return true
	}

	nbDrives = remainingDistance / c.speed

	if remainingDistance%c.speed > 0 {
		nbDrives += 1
	}

	return c.batteryDrain*nbDrives <= c.battery
}
