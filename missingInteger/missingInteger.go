package missinginteger

func Solution(A []int) int {
    mapOfA := make(map[int]struct{})
    for _, item := range A {
        mapOfA[item] = struct{}{}
    }
    for i := 1; i < 100_000_000; i++ {
        if _, exists := mapOfA[i]; !exists {
            return i
        }
    }
    return 1
}