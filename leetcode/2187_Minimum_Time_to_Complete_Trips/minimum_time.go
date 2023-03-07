package main

func minimumTime(time []int, totalTrips int) int64 {
	maxBusTime := 0
	for i := 0; i < len(time); i++ {
		if time[i] > maxBusTime {
			maxBusTime = time[i]
		}
	}

	calculatedTotalTrip := 0

	for i := 0; i < len(time); i++ {
		calculatedTotalTrip += maxBusTime / time[i]
	}

	if calculatedTotalTrip == totalTrips {
		return int64(maxBusTime)
	}

	left := 0
	right := maxBusTime

	if calculatedTotalTrip < totalTrips {
		countRings := totalTrips / calculatedTotalTrip
		right = maxBusTime * (countRings + 1)
	}

	for left < right {
		mid := (left + right) >> 1
		if check(time, totalTrips, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}

func check(time []int, totalTrips, mid int) bool {
	cnt := 0
	for _, t := range time {
		cnt += mid / t
		if cnt >= totalTrips {
			return true
		}
	}
	return false
}
