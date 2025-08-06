package main

import (
	"fmt"
	"time"

	"main.go/patterngen"
)

func main() {
	pattern := patterngen.GenerateDailyPattern(0, 1440, 1000.0)
	for minute := 0; minute < 1440; minute++ {
		value := pattern[minute]
		fmt.Printf("%02d:%02d\t%.1f\n", minute/60, minute%60, value)
		time.Sleep(10 * time.Millisecond) // Hızlı simülasyon
	}
}
