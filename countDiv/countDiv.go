package countdiv

import "math"

func Solution(A int, B int, K int) int {
	if K == 1 {
		return A - B + 1
	}
	minDivAfterA := int64(math.Ceil(float64(A)/float64(K)) * float64(K))
	maxDivBeforeB := int64(math.Floor(float64(B)/float64(K)) * float64(K))
	noOfDeviablesByK := ((maxDivBeforeB - minDivAfterA) / int64(K)) + 1
	return int(noOfDeviablesByK)
}
