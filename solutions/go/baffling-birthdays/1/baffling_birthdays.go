package bafflingbirthdays

import (
	"math/rand"
	"time"
)

func SharedBirthday(dates []time.Time) bool {
	for i := 0; i < len(dates); i++ {
		for j := i + 1; j < len(dates); j++ {
			if dates[i].Month() == dates[j].Month() {
				if dates[i].Day() == dates[j].Day() {
					return true
				}
			}
		}
	}
	return false
}

func RandomBirthdates(size int) []time.Time {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))
	birthdates := make([]time.Time, size)

	for i := 0; i < size; i++ {
		n := r.Intn(365)
		var t time.Time
		t.AddDate(1970, 1, 1)
		t = t.Add(time.Duration(86400000000000 * n))
		birthdates[i] = t
	}
	return birthdates
}

func EstimatedProbability(size int) float64 {
	if size <= 1 {
		return 0.0
	}

	const simulations = 1000
	count := 0
	for i := 0; i < simulations; i++ {
		if SharedBirthday(RandomBirthdates(size)) {
			count++
		}
	}
	return float64(count * 100.0 / simulations)
}
