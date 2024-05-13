package tapeequilibrium

func Solution(A []int) int {
	if len(A) <= 1 {
		return 1
	}
	total := 0
	sumList := make([]int, len(A))
	for i, item := range A {
		total += item
		if i == 0 {
			sumList[i] = item
		} else {
			sumList[i] = sumList[i - 1] + item
		}
	}
	minDifference := 2147483647
	var sumArray0, sumArray1 int
	for p := 1; p < len(A); p++ {
		sumArray0 = sumList[p - 1]
		sumArray1 = total - sumArray0
		if diff := diff(sumArray0, sumArray1); diff < minDifference {
			minDifference = diff
		}
	}
	return minDifference
}

func diff(val0, val1 int) int {
	if val0 >= val1 {
		return val0 - val1
	}
	return val1 - val0
}
