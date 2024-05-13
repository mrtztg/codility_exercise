package cyclicrotation

func Solution(A []int, K int) []int {
	if K >= len(A) {
		K -= len(A)
	} else if K == len(A) {
		return A
	}
	finalArr := append(A[len(A) - K:], A[:len(A) - K]...)
	return finalArr
}