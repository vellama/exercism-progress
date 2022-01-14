package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var choice string

	if option1 < option2 {
		choice = option1
	} else {
		choice = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", choice)
}

func computePrice(originalPrice, percentageOff float64) float64 {
	return originalPrice * percentageOff / 100
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return computePrice(originalPrice, 80)
	} else if age >= 10 {
		return computePrice(originalPrice, 50)
	} else {
		return computePrice(originalPrice, 70)
	}
}
