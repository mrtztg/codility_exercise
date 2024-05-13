package tapeequilibrium

import "math"

func Solution(A []int) int{
	if len(A) <= 1 {
		return 1
	}
	minDifference := math.MaxInt 
	for p := 1; p < len(A); p++ {
		sumArray0, sumArray1 := sum(A[:p]), sum(A[p:])
		if diff := diff(sumArray0,sumArray1); diff < minDifference {
			minDifference = diff
		}
	}
	return minDifference
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
func diff(val0, val1 int) int {
	if val0 >= val1 {
		return val0 - val1
	}
	return val1 - val0
}