package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lnums1 := len(nums1)
	lnums2 := len(nums2)

	if lnums1 == 0 {
		return getMedian(nums2)
	}
	if lnums2 == 0 {
		return getMedian(nums1)
	}

	l := lnums1 + lnums2

	arr := make([]int, l)

	j := 0
	for i := 0; i < lnums1 || i < lnums2; i++ {
		if i < lnums1 {
			arr[j] = nums1[i]
			j++
		}
		if i < lnums2 {
			arr[j] = nums2[i]
			j++
		}
		d := j
		for d > 1 {
			fmt.Println(arr)
			if arr[d-1] < arr[d-2] {
				arr[d-1], arr[d-2] = arr[d-2], arr[d-1]
			}
			d--
		}

	}

	return getMedian(arr)
}

func getMedian(arr []int) float64 {
	l := len(arr)
	fmt.Println(arr)
	if l%2 == 0 {
		return float64(arr[l/2]+arr[l/2-1]) / 2
	}
	return float64(arr[l/2])
}

func main() {
	fmt.Println(3 / 2)
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
