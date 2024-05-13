package permmissingelem

func Solution(A []int) int {
	if len(A) < 2 {
		return 0
	}
	sortedArr := make([]bool, len(A) + 2)
	for _, item := range A {
		sortedArr[item - 1] = true
	}
	for i, exists := range sortedArr {
		if !exists {
			return i + 1
		}
	}
	return 0
}