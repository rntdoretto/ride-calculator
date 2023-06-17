package cognitivecomplexityrefactoring_test

import (
	"testing"
	"time"

	cognitivecomplexityrefactoring "github.com/rntdoretto/ride-calculator/04-cognitive-complexity-refactoring"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name     string
	input    []cognitivecomplexityrefactoring.Segment
	expected float64
}

func TestExecute(t *testing.T) {
	testTable := []testTable{
		{
			name: "Should calculate the fare for a minimal ride",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 2.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 10.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 21.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day with two segments",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 42.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day overnight",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 39.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: 29.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday overnight",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 18, 23, 0, 0, 0, time.UTC),
				},
			},
			expected: 50.0,
		},
		{
			name: "Should return -1 if distanceance is invalid",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: -10.0,
					DateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expected: -1,
		},
		{
			name: "Should return -2 if date is invalid",
			input: []cognitivecomplexityrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expected: -2,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fare := cognitivecomplexityrefactoring.CalculateRide(tt.input)
			assert.Equal(t, tt.expected, fare)
		})
	}
}
