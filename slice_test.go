package nombre

import (
	"fmt"
	"testing"
)

func TestContains(t *testing.T) {
	type testCase[T any] struct {
		slice    []T
		elem     T
		expected bool
	}

	testCasesRune := []testCase[rune]{
		{[]rune{'a', 'b', 'c'}, 'c', true},
		{[]rune{'a', 'b', 'c'}, 'd', false},
	}

	testCasesInt := []testCase[int]{
		{[]int{}, 3, false},
		{[]int{2, 6, 4, 98}, 3, false},
		{[]int{2, 6, 4, 98}, 2, true},
	}

	t.Run("first letter of empty string", func(t *testing.T) {
		empty := ""
		vowels := []rune{'a', 'e', 'i', 'o', 'u'}
		if contains(vowels[:], firstLetter(empty)) {
			t.Fail()
		}
	})

	for _, tc := range testCasesRune {
		desc := fmt.Sprint(tc.slice)

		t.Run(desc, func(t *testing.T) {
			got := contains(tc.slice, tc.elem)
			if tc.expected != got {
				t.Fail()
			}
		})
	}

	for _, tc := range testCasesInt {
		desc := fmt.Sprint(tc.slice)

		t.Run(desc, func(t *testing.T) {
			got := contains(tc.slice, tc.elem)
			if tc.expected != got {
				t.Fail()
			}
		})
	}
}
