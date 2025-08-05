package main

import (
	"math"
	"math/rand/v2"
	"time"
)

func GenerateDailyPattern(startMinute, endMinute int, maxValue float64) []float64 {
	totalMinutes := endMinute - startMinute
	result := make([]float64, totalMinutes)

	for i := 0; i < totalMinutes; i++ {
		minuteOfDay := startMinute + i
		sinVal := math.Sin(math.Pi * float64(minuteOfDay-startMinute) / float64(endMinute-startMinute))
		value := maxValue * sinVal

		var noise float64
		// Öğlen (11:00 - 15:00) arası daha yüksek gürültü
		if minuteOfDay >= 660 && minuteOfDay <= 900 {
			noise = rand.NormFloat64() * maxValue * 0.09
		} else {
			noise = rand.NormFloat64() * maxValue * 0.02
		}
		value += noise

		if value < 0 {
			value = 0
		}
		result[i] = value
	}
	return result
}

func getCurrentMinuteOfDay() int {
	now := time.Now()
	return now.Hour()*60 + now.Minute()
}

func getPatternValueForNow(pattern []float64, startMinute, endMinute int) float64 {
	minuteOfDay := getCurrentMinuteOfDay()
	if minuteOfDay < startMinute || minuteOfDay >= endMinute {
		return 0
	}
	return pattern[minuteOfDay-startMinute]
}
