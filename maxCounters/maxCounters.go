package maxcounters

func Solution(N int, A []int) []int {
    arr := make([]int, N)
    maxEncounter := 0
    maxNo := 0
    for _, item := range A {
        if (item <= N) {
            destPos := item - 1
            if arr[destPos] < maxEncounter {
                arr[destPos] = maxEncounter
            }
            arr[destPos] += 1
            if arr[destPos] > maxNo {
                maxNo = arr[destPos]
            }
        } else {
            maxEncounter = maxNo
        }
    }
    for i, item := range arr {
        if item < maxEncounter {
            arr[i] = maxEncounter
        }
    }
    return arr
}