package algo

func bynarySearch(arr []int, n int) (int, bool) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == n {
			return mid, true
		}

		if n < arr[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return 0, false
}
