package nombre

type Options struct {
	spaceAfterMinus              bool
	spaceBetweenOrderOfMagnitude bool
}

func DefaultOptions() Options {
	return Options{
		spaceAfterMinus:              false,
		spaceBetweenOrderOfMagnitude: false,
	}
}
