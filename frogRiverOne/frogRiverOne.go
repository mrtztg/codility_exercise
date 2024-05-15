package frogriverone

func Solution(X int, A []int) int {
    if X > len(A) {
        return -1
    }
    filledSteps := make([]int, X)
    for second, position := range A {
        if filledSteps[position - 1] == 0 {
            filledSteps[position - 1] = second + 1
        }
    }
    secondOfLatestStep := 0
    for _, second := range filledSteps {
        if second == 0 {
            return -1
        }
        if second >= secondOfLatestStep {
			secondOfLatestStep = second
        }
    }
    return secondOfLatestStep - 1
}
