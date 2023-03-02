package main

// impl merg_sort from https://en.wikipedia.org/wiki/Merge_sort
func sortArray(nums []int) []int {
	l := len(nums)
	if l == 0 || l == 1 {
		return nums
	}
	mid := l / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])

	m := 0
	n := 0
	k := 0

	tmp := make([]int, len(left)+len(right))

	for n < len(left) && m < len(right) {
		if left[n] <= right[m] {
			tmp[k] = left[n]
			n++
		} else {
			tmp[k] = right[m]
			m++
		}
		k++
	}
	for n < len(left) {
		tmp[k] = left[n]
		n++
		k++
	}
	for m < len(right) {
		tmp[k] = left[m]
		m++
		k++
	}

	return tmp
}
