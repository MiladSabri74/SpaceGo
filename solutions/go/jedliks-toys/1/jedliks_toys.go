package jedlik

import "fmt"

func (c *Car) Drive() {
	if c.battery-c.batteryDrain >= 0 {
		c.distance += c.speed
		c.battery -= c.batteryDrain
	}
}

func (c Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c Car) CanFinish(trackDistance int) bool {
	for c.distance < trackDistance {
		if c.battery-c.batteryDrain >= 0 {
			c.distance += c.speed
			c.battery -= c.batteryDrain
		} else {
			return false
		}
	}
	return true
}
