package main

// first try solution
func reverse_integer(x int) int {
	xprime := 0
	sign := 1;
	if x < 0 {
		sign = -1
		x = -x
	}
	for x > 0 { 
		xprime = xprime * 10 + x % 10
		x /= 10
	}
	return xprime * sign
}

// better and with limitations
func reverse(x int) int {
    var xprime int
    const MaxInt = int(^uint32(0) >> 1)
	
	for x != 0 {
		xprime = xprime * 10 + x % 10
		x /= 10
        if xprime > MaxInt || xprime < -MaxInt - 1 {
            return 0
        }
	}

	return xprime
}
