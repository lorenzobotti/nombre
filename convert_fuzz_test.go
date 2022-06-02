package nombre

import (
	"strings"
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

	f.Fuzz(func(t *testing.T, a int) {
		got := Convert(a)
		if strings.Contains(got, " ") {
			t.Fail()
		}
	})
}
