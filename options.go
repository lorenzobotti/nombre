package nombre

type Options struct {
	// Put a space after "meno"
	SpaceAfterMinus bool
	// Put a space between orders of magnitude.
	// Example: "un miliardo novecentotrentadue milioni"
	SpaceBetweenOrderOfMagnitude bool
	// Shorten the "one" before the name of an order
	// of magnitude.
	// Example: "centoventunmila quattrocentosettantadue"
	ShortenOne bool
	// Put an accent on "tre", but not when it's on
	// its own (including when orders of magnitude are
	// spaced out)
	AccentThree bool
}

func DefaultOptions() Options {
	return Options{
		SpaceAfterMinus:              false,
		SpaceBetweenOrderOfMagnitude: false,
		ShortenOne:                   false,
		AccentThree:                  false,
	}
}
