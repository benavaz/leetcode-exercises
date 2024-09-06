func divide(dividend int, divisor int) int {
	dividendSign := 1
	divisorSign := 1
	sign := -1
	if dividend+dividend < 0 {
		dividendSign = -1
		dividend = dividend - dividend - dividend
	}

	if divisor+divisor < 0 {
		divisorSign = -1
		divisor = divisor - divisor - divisor
	}

	if (dividendSign < 0 && divisorSign < 0) || (dividendSign > 0 && divisorSign > 0) {
		sign = 1
	}

	if divisor == 1 {
		if sign == 1 && dividend == 2147483648 {
			return dividend - 1
		}

		if sign > 0 {
			return dividend
		} else {
			return dividend - dividend - dividend
		}
	}

	// solution with recursion
	// return sign * helper(dividend, divisor)

	count := 0
	for dividend >= divisor {
		count += 1
		dividend = dividend - divisor
	}
	if sign > 0 {
		return count
	} else {
		return count - count - count
	}
}

// this can also be solved using recursion
// the problem is that when the dividend is very big
// and the divisor very small, this results
// in many recursive calls and exceeds time limit
func helper(dividend int, divisor int) int {

	if dividend < divisor {
		return 0
	} else {
		return 1 + divide(dividend-divisor, divisor)
	}

}