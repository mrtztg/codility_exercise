package binarygap

import "fmt"

func Solution(N int) int {
	b := fmt.Sprintf("%b", N)
	currentCount := 0
	biggestProven := 0
	var one rune = '1'
	for _, rn := range b {
		if rn == one {
			if currentCount > biggestProven {
				biggestProven = currentCount
			}
			currentCount = 0
		} else {
			currentCount++
		}
	}
	return biggestProven
}
