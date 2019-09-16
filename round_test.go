package rounding

import (
	"fmt"
	"testing"
)

func TestRoundDown(t *testing.T) {
	testCases := []struct {
		value         float64
		decimalPlaces int
		expected      float64
	}{
		{
			value:         12,
			decimalPlaces: 1,
			expected:      12,
		},
		{
			value:         15.1,
			decimalPlaces: 0,
			expected:      15,
		},
		{
			value:         15.1,
			decimalPlaces: 1,
			expected:      15.1,
		},
		{
			value:         18.14,
			decimalPlaces: 1,
			expected:      18.1,
		},
		{
			value:         25.18,
			decimalPlaces: 1,
			expected:      25.1,
		},
		{
			value:         115.18123,
			decimalPlaces: 1,
			expected:      115.1,
		},
		{
			value:         115.18123,
			decimalPlaces: 2,
			expected:      115.18,
		},
		{
			value:         115.18123,
			decimalPlaces: 4,
			expected:      115.1812,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v - %v", tc.value, tc.decimalPlaces), func(t *testing.T) {
			if got := RoundDown(tc.value, tc.decimalPlaces); got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestRoundUp(t *testing.T) {
	testCases := []struct {
		value         float64
		decimalPlaces int
		expected      float64
	}{
		{
			value:         12,
			decimalPlaces: 1,
			expected:      12,
		},
		{
			value:         15.1,
			decimalPlaces: 0,
			expected:      16,
		},
		{
			value:         15.1,
			decimalPlaces: 1,
			expected:      15.1,
		},
		{
			value:         18.14,
			decimalPlaces: 1,
			expected:      18.2,
		},
		{
			value:         25.18,
			decimalPlaces: 1,
			expected:      25.2,
		},
		{
			value:         115.18123,
			decimalPlaces: 1,
			expected:      115.2,
		},
		{
			value:         115.18123,
			decimalPlaces: 2,
			expected:      115.19,
		},
		{
			value:         115.18123,
			decimalPlaces: 4,
			expected:      115.1813,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v - %v", tc.value, tc.decimalPlaces), func(t *testing.T) {
			if got := RoundUp(tc.value, tc.decimalPlaces); got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func TestRoundNearest(t *testing.T) {
	testCases := []struct {
		value         float64
		decimalPlaces int
		expected      float64
	}{
		{
			value:         12,
			decimalPlaces: 1,
			expected:      12,
		},
		{
			value:         15.1,
			decimalPlaces: 0,
			expected:      15,
		},
		{
			value:         15.1,
			decimalPlaces: 1,
			expected:      15.1,
		},
		{
			value:         18.14,
			decimalPlaces: 1,
			expected:      18.1,
		},
		{
			value:         25.18,
			decimalPlaces: 1,
			expected:      25.2,
		},
		{
			value:         115.18123,
			decimalPlaces: 1,
			expected:      115.2,
		},
		{
			value:         115.18123,
			decimalPlaces: 2,
			expected:      115.18,
		},
		{
			value:         115.18123,
			decimalPlaces: 4,
			expected:      115.1812,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v - %v", tc.value, tc.decimalPlaces), func(t *testing.T) {
			if got := RoundNearest(tc.value, tc.decimalPlaces); got != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}
