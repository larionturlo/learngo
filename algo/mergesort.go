package algo

func merge(a, b []int) (r []int) {
	lenA := len(a)
	lenB := len(b)
	if lenA == 0 {
		r = b
	} else if lenB == 0 {
		r = a
	} else {
		i := 0
		j := 0

		for i < lenA || j < lenB {
			if j > lenB-1 {
				r = append(r, a[i])
				i++
				continue
			}

			if i > lenA-1 {
				r = append(r, b[j])
				j++
				continue
			}

			if a[i] < b[j] {
				r = append(r, a[i])
				i++
			} else {
				r = append(r, b[j])
				j++
			}
		}
		if lenA-1 <= i {
			r = append(r, a[i:]...)
		}
		if lenB-1 <= j {
			r = append(r, b[j:]...)
		}
	}
	return
}

func mergeSort(arr []int) []int {
	for j := 2; j/2 <= len(arr); j = j * 2 {
		for i := 0; i < len(arr); i = i + j {
			d := j / 2

			mid := i + d
			right := i + j

			if mid > len(arr) {
				mid = len(arr)
			}

			if right > len(arr) {
				right = len(arr)
			}

			mrg := merge(arr[i:mid], arr[mid:right])
			arr = append(arr[:i], append(mrg, arr[right:]...)...)
		}
	}
	return arr
}
