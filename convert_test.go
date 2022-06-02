package nombre

import (
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{0, "zero"},
		{1, "uno"},
		{-1, "menouno"},
		{30, "trenta"},
		{42, "quarantadue"},
		{51, "cinquantuno"},
		{18, "diciotto"},
		{101, "centouno"},
		{151, "centocinquantuno"},
		{1452, "millequattrocentocinquantadue"},
		{1984, "millenovecentoottantaquattro"},
		{2002, "duemiladue"},
		{400002, "quattrocentomiladue"},
		{182952, "centoottantaduemilanovecentocinquantadue"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			got := Convert(tc.input)

			if got != tc.expected {
				t.Fatalf("expected '%s', got '%s'", tc.expected, got)
			}
		})
	}
}

func TestConvertWithOptions(t *testing.T) {
	testCases := []struct {
		input                        int
		spaceAfterMinus              bool
		spaceBetweenOrderOfMagnitude bool
		expected                     string
	}{
		{0, true, true, "zero"},
		{1, true, true, "uno"},
		{-1, true, true, "meno uno"},
		{30, true, true, "trenta"},
		{42, true, true, "quarantadue"},
		{51, true, true, "cinquantuno"},
		{18, true, true, "diciotto"},
		{101, true, true, "centouno"},
		{-151, true, true, "meno centocinquantuno"},
		{1452, true, true, "mille quattrocentocinquantadue"},
		{1984, true, true, "mille novecentoottantaquattro"},
		{-2002, true, true, "meno duemila due"},
		{400002, true, true, "quattrocentomila due"},
		{182952, true, true, "centoottantaduemila novecentocinquantadue"},
	}

	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			options := Options{
				SpaceAfterMinus:              tc.spaceAfterMinus,
				SpaceBetweenOrderOfMagnitude: tc.spaceBetweenOrderOfMagnitude,
			}
			got := ConvertWithOptions(tc.input, options)

			if got != tc.expected {
				t.Fatalf("expected '%s', got '%s'", tc.expected, got)
			}
		})
	}
}

func TestGetDigit(t *testing.T) {
	testCases := []struct {
		number   int
		digit    int
		expected int
	}{
		{100, 1, 0},
		{100, 2, 0},
		{100, 3, 1},
		{87, 2, 8},
		{1765, 3, 7},
		{1986473, 6, 9},
	}
	for _, tc := range testCases {
		str := strconv.Itoa(tc.number)
		t.Run(str, func(t *testing.T) {
			got := getDigit(tc.number, tc.digit)
			if tc.expected != got {
				t.Fatalf("%d, digit %d, expected %d, got %d", tc.number, tc.digit, tc.expected, got)
			}
		})
	}
}

func TestGetDigits(t *testing.T) {
	testCases := []struct {
		number     int
		start, end int
		expected   int
	}{
		{1234, 3, 2, 23},
		{1874, 3, 2, 87},
		{2874764, 5, 2, 7476},
		{2874764, 5, 3, 747},
		{10000001, 5, 3, 0},
		{10000101, 5, 3, 1},
		{10010101, 5, 3, 101},
		{1, 6, 4, 0},
	}
	for _, tc := range testCases {
		desc := strconv.Itoa(tc.number)
		t.Run(desc, func(t *testing.T) {
			got := getDigits(tc.number, tc.start, tc.end)
			if tc.expected != got {
				t.Fatalf(
					"%d, start %d end %d, expected %d, got %d",
					tc.number,
					tc.start,
					tc.end,
					tc.expected,
					got,
				)
			}
		})
	}
}
