package nombre

import "strings"

func Convert(number int) string {
	if number == 0 {
		return "zero"
	}

	negative := false
	if number < 0 {
		negative = true
		number *= -1
	}

	hundreds := formatHundreds(getDigit(number, 3))
	tens := formatTens(number % 100)

	out := strings.Builder{}
	maxOrder := maxOrderOfMagnitude(number)
	biggestOrderInThrees := maxOrder + (3 - (maxOrder % 3))

	if negative {
		out.WriteString("meno")
	}

	for order := biggestOrderInThrees; order >= 3; order -= 3 {
		out.WriteString(formatOrderOfMagnitude(number, order))
	}

	out.WriteString(hundreds)
	out.WriteString(tens)

	return out.String()
}

func formatTens(ts int) string {
	t := getDigit(ts, 2)
	o := getDigit(ts, 1)

	if t == 1 {
		return teens[o]
	} else {
		tensStr := tens[t]
		onesStr := digits[o]

		if contains(vowels[:], firstLetter(onesStr)) {
			if len(tensStr) > 0 {
				tensStr = tensStr[:len(tensStr)-1]
			}
		}

		return tensStr + onesStr
	}
}

func formatHundreds(digit int) string {
	if digit == 0 {
		return ""
	}

	if digit == 1 {
		return "cento"
	}

	return digits[digit] + "cento"
}

func formatOrderOfMagnitude(number, order int) string {
	terms, areThereTerms := ordersOfMagnitude[order]
	if !areThereTerms {
		return ""
	}
	singular := terms[0]
	plural := terms[1]

	digits := getOrderOfMagnitude(number, order)
	if digits == 0 {
		return ""
	}
	if digits == 1 {
		return singular
	}

	return Convert(digits) + plural
}
