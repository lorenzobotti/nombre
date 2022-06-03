package nombre

import (
	"math/rand"
	"testing"
)

func FuzzConvert(f *testing.F) {
	f.Add(476287)
	f.Add(98954)
	f.Add(98)
	f.Add(1)
	f.Add(0)
	f.Add(8572702459270)
	f.Add(7895800005278)
	f.Add(70000000)
	f.Add(50000000)

	f.Fuzz(func(t *testing.T, num int) {
		got := ConvertWithOptions(num, randomOptions())
		if len(got) == 0 {
			t.Fail()
		}
	})
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomOptions() Options {
	return Options{
		randomBool(),
		randomBool(),
		randomBool(),
		randomBool(),
	}
}
