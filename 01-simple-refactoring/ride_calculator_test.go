package simplerefactoring_test

import (
	"testing"
	"time"

	simplerefactoring "github.com/rntdoretto/ride-calculator/01-simple-refactoring"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name     string
	input    []map[string]any
	expected float64
}

func TestExecute(t *testing.T) {
	testTable := []testTable{
		{
			name: "Should calculate the fare for a minimal ride",
			input: []map[string]any{
				{
					"distance": 2.0,
					"dateTime": time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 10.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 21.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day with two segments",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 42.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day overnight",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 17, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 39.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 29.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday overnight",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": time.Date(2023, 6, 18, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 50.0,
		},
		{
			name: "Should return -1 if distanceance is invalid",
			input: []map[string]any{
				{
					"distance": -10.0,
					"dateTime": time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: -1,
		},
		{
			name: "Should return -2 if date is invalid",
			input: []map[string]any{
				{
					"distance": 10.0,
					"dateTime": "asdf",
				},
			},
			expected: -2,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fare := simplerefactoring.CalculateRide(tt.input)
			assert.Equal(t, tt.expected, fare)
		})
	}
}
