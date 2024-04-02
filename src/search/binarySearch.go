package search

func BinarySearch(arr []int, target int) int {
	i, f := 0, len(arr)-1

	for i >= f {
		m := (i + f) / 2
		if target == arr[m] {
			return m
		} else if target > arr[m] {
			i = m + 1
		} else {
			f = m - 1
		}
	}

	return -1
}
