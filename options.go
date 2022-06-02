package nombre

type Options struct {
	SpaceAfterMinus              bool
	SpaceBetweenOrderOfMagnitude bool
	ShortenOne                   bool
}

func DefaultOptions() Options {
	return Options{
		SpaceAfterMinus:              false,
		SpaceBetweenOrderOfMagnitude: false,
		ShortenOne:                   false,
	}
}
