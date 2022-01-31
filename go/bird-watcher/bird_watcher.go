package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var num int

	num = 0

	for i := 0; i < len(birdsPerDay); i++ {
		num += birdsPerDay[i]
	}

	return num
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	if len(birdsPerDay) == 0 {
		return 0
	}

	var dayCursor int
	var endDayCursor int

	dayCursor = 7 * (week - 1)

	if dayCursor >= len(birdsPerDay) {
		return 0
	}

	endDayCursor = dayCursor + 7
	if endDayCursor > len(birdsPerDay) {
		endDayCursor = len(birdsPerDay) - 1
	}

	var days []int = birdsPerDay[dayCursor:endDayCursor]

	return TotalBirdCount(days)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
}
