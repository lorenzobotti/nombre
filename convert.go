package nombre

import (
	"strings"
)

// Convert turns a number into a string using the default options
func Convert(number int) string {
	return ConvertWithOptions(number, DefaultOptions())
}

// ConvertWithOptions turns a number into a string using the provided options
func ConvertWithOptions(number int, options Options) string {
	if number == 0 {
		return "zero"
	}

	negative := false
	if number < 0 {
		negative = true
		number *= -1
	}

	out := strings.Builder{}
	maxOrder := maxOrderOfMagnitude(number)
	biggestOrderInThrees := maxOrder + (3 - (maxOrder % 3))

	const wroteNothing = 0
	const wroteMinus = 1
	const wroteNumber = 2

	lastWritten := wroteNothing
	shouldWriteSpace := func() bool {
		return (lastWritten == wroteNumber && options.SpaceBetweenOrderOfMagnitude) ||
			(lastWritten == wroteMinus && options.SpaceAfterMinus)
	}

	if negative {
		out.WriteString("meno")
		lastWritten = wroteMinus
	}

	for order := biggestOrderInThrees; order >= 3; order -= 3 {
		wroteSomething := formatOrderOfMagnitude(number, order, options.ShortenOne, shouldWriteSpace(), options.SpaceBetweenOrderOfMagnitude, options.AccentThree, &out)

		if wroteSomething {
			lastWritten = wroteNumber
		}
	}

	accentThree := (lastWritten == wroteNumber) && options.AccentThree
	formatFirstThree(number%1000, false, shouldWriteSpace(), accentThree, &out)

	return out.String()
}

// formatFirstThree formats numbers lesser than one thousand.
func formatFirstThree(ts int, shortenOne, shouldStartWithSpace, accentThree bool, builder *strings.Builder) (wroteSomething bool) {
	haveWrittenSpace := false
	writeString := func(strs ...string) {
		if shouldStartWithSpace && !haveWrittenSpace {
			builder.WriteRune(' ')
			haveWrittenSpace = true
		}

		for _, s := range strs {
			builder.WriteString(s)
		}
	}

	lengthBefore := builder.Len()

	h := getDigit(ts, 3)
	t := getDigit(ts, 2)
	o := getDigit(ts, 1)

	if h == 1 {
		writeString("cento")
	} else if h != 0 {
		writeString(digits[h], "cento")
	}

	if t == 1 {
		writeString(teens[o])
	} else {
		tensStr := tens[t]
		onesStr := digits[o]

		if t != 0 && h != 0 && o == 3 && accentThree {
			onesStr = "trè"
		}

		// shorten the last letter of the tens if the unit starts with a vowel
		// example: "quarantauno" -> "quarantuno"
		if contains(vowels[:], firstLetter(onesStr)) {
			if len(tensStr) > 0 {
				tensStr = tensStr[:len(tensStr)-1]
			}
		}

		if o == 1 && shortenOne {
			writeString(tensStr, "un")
		} else {
			writeString(tensStr, onesStr)
		}
	}

	return lengthBefore != builder.Len()
}

// formatOrderOfMagnitude picks an order of magnitude from a number and formats only that,
// together with the name of the order of magnitude. The names are described by the global
// variable `ordersOfMagnitude`
func formatOrderOfMagnitude(number, order int, shortenOne, startWithSpace, spaceBetween, accentThree bool, builder *strings.Builder) (wroteSomething bool) {
	lengthBefore := builder.Len()
	didIWriteSomething := func() bool { return lengthBefore != builder.Len() }

	terms, areThereTerms := ordersOfMagnitude[order]
	if !areThereTerms {
		return didIWriteSomething()
	}

	singular := terms[0]
	plural := terms[1]

	digits := getOrderOfMagnitude(number, order)
	if digits == 1 {
		if startWithSpace {
			builder.WriteRune(' ')
		}
		builder.WriteString(singular)
	} else if digits != 0 {
		accentThree = accentThree && spaceBetween

		formatFirstThree(digits, shortenOne, startWithSpace, accentThree, builder)
		if spaceBetween && order != 3 {
			builder.WriteRune(' ')
		}
		builder.WriteString(plural)
	}

	return didIWriteSomething()
}
