package nombre

import (
	"strconv"
	"testing"
)

func TestGetMaxOrder(t *testing.T) {
	testCases := []struct {
		number   int
		expected int
	}{
		{12, 2},
		{121, 3},
		{1213, 4},
		{9999, 4},
		{10000, 5},
	}
	for _, tc := range testCases {
		desc := strconv.Itoa(tc.number)
		t.Run(desc, func(t *testing.T) {
			got := maxOrderOfMagnitude(tc.number)
			if tc.expected != got {
				t.Fatalf("%d, expected %d, got %d", tc.number, tc.expected, got)
			}
		})
	}
}

func TestGetOrder(t *testing.T) {
	testCases := []struct {
		number   int
		order    int
		expected int
	}{
		{123654789, 3, 654},
		{654789, 3, 654},
		{654000, 3, 654},
		{876654000, 3, 654},
		{876000000, 3, 0},
		{876000001, 0, 1},
	}
	for _, tc := range testCases {
		desc := strconv.Itoa(tc.number)
		t.Run(desc, func(t *testing.T) {
			got := getOrderOfMagnitude(tc.number, tc.order)
			if tc.expected != got {
				t.Fatalf("%d order %d: expected %d, got %d", tc.number, tc.order, tc.expected, got)
			}
		})
	}
}
