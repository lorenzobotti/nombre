package nombre

func getDigit(number, fromTheRight int) int {
	orderOfMagnitude := pow(10, fromTheRight-1)
	noBigger := number % (orderOfMagnitude * 10)
	digit := noBigger / orderOfMagnitude

	return digit
}

func getDigits(number, start, end int) int {
	if start < end {
		panic("start must be less than end!")
	}

	startMagnitude := pow(10, start-1)
	endMagnitude := pow(10, end-1)

	noBigger := number % (startMagnitude * 10)
	onlyDigits := noBigger / endMagnitude

	return onlyDigits
}

func getOrderOfMagnitude(number, order int) int {
	return getDigits(number, order+3, order+1)
}

func maxOrderOfMagnitude(number int) int {
	order, tens := 1, 10
	for {
		if number < tens {
			return order
		}

		order += 1
		tens *= 10
	}
}

func pow(base, exp int) int {
	if exp == 0 {
		if base == 0 {
			panic("zero to the power of zero")
		} else {
			return 1
		}
	}

	res := base
	for i := 1; i < exp; i++ {
		res *= base
	}

	return res
}
