package nombre

type Options struct {
	spaceAfterMinus              bool
	spaceBetweenOrderOfMagnitude bool
	shortenOne                   bool
}

func DefaultOptions() Options {
	return Options{
		spaceAfterMinus:              false,
		spaceBetweenOrderOfMagnitude: false,
		shortenOne:                   false,
	}
}
