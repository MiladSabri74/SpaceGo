package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0
	for _, birds := range birdsPerDay {
		totalBirds += birds
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalBirds := 0
	for _, birds := range birdsPerDay[(week-1)*7 : week*7] {
		totalBirds += birds
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for day, _ := range birdsPerDay {
		if day%2 == 0 {
			birdsPerDay[day] = birdsPerDay[day] + 1
		}
	}
	return birdsPerDay
}
