package frogjmp

import "math"

func Solution(X int, Y int, D int) int {
	if X == Y {
		return 0
	}
	distance := Y - X
	return int(math.Ceil(float64(distance) / float64(D)))
}
