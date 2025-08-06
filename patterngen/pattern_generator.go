package patterngen

import (
	"math/rand/v2"
	"time"
)

// Günlük bir desen oluşturur
// startMinute: Desenin başlangıç dakikası (0-1439 arası)
// endMinute: Desenin bitiş dakikası (0-1439 arası, startMinute'dan büyük olmalı)
func GenerateDailyPattern(startMinute, endMinute int, maxValue float64) []float64 {
	pattern := make([]float64, endMinute-startMinute)
	for minute := startMinute; minute < endMinute; minute++ {
		hour := float64(minute) / 60.0
		var value float64

		switch {
		case hour < 6.0: // Gece
			value = 0
		case hour < 8.0: // Güneş doğuşu (yavaş artış)
			value = maxValue * 0.2 * (hour - 6.0) / 2.0
		case hour < 11.0: // Sabah hızlı artış
			value = maxValue * (0.2 + 0.8*(hour-8.0)/3.0)
		case hour < 15.0: // Öğlen (peak, noise ile)
			base := maxValue
			noise := rand.NormFloat64() * maxValue * 0.2 // %20 oynaklık
			value = base + noise
		case hour < 18.0: // Akşam yavaş azalış
			value = maxValue * (1.0 - (hour-15.0)/3.0)
		default: // Gece
			value = 0
		}

		if value < 0 {
			value = 0
		}
		pattern[minute-startMinute] = value
	}
	return pattern
}

// Şu anki günün dakikasını döner (0-1439 arası)
// 0 = 00:00, 1439 = 23:59
// Bu fonksiyon, günün saatini ve dakikasını kullanarak toplam dakikayı hesaplar.
func GetCurrentMinuteOfDay() int {
	now := time.Now()
	return now.Hour()*60 + now.Minute()
}

func GetPatternValueForNow(pattern []float64, startMinute, endMinute int) float64 {
	minuteOfDay := GetCurrentMinuteOfDay()
	if minuteOfDay < startMinute || minuteOfDay >= endMinute {
		return 0
	}
	return pattern[minuteOfDay-startMinute]
}
