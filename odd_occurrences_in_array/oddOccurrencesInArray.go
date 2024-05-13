package oddoccurrencesinarray

func Solution(A []int) int {
	oddValues := map[int]struct{}{}
	for _, item := range A {
		if _, ok := oddValues[item]; ok {
			delete(oddValues, item)
		} else {
			oddValues[item] = struct{}{}
		}
	}
	for k := range oddValues {
		return k
	}
	return 0
}
