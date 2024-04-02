package search

func LinearSearch(arr []int, target int) int {
	for i := 0; i < len(arr)-1; i++ {
		if target == arr[i] {
			return i
		}
	}

	return -1
}
