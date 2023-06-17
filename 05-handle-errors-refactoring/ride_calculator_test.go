package handleerrorsrefactoring_test

import (
	"testing"
	"time"

	handleerrorsrefactoring "github.com/rntdoretto/ride-calculator/05-handle-errors-refactoring"
	"github.com/stretchr/testify/assert"
)

type testTable struct {
	name         string
	input        []handleerrorsrefactoring.Segment
	expectedFare float64
	expectedErr  string
}

func TestExecute(t *testing.T) {
	testTable := []testTable{
		{
			name: "Should calculate the fare for a minimal ride",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 2.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 10.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 21.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day with two segments",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 42.0,
		},
		{
			name: "Should calculate the fare for a ride on a normal day overnight",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 17, 23, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 39.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 29.0,
		},
		{
			name: "Should calculate the fare for a ride on a sunday overnight",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(2023, 6, 18, 23, 0, 0, 0, time.UTC),
				},
			},
			expectedFare: 50.0,
		},
		{
			name: "Should return -1 if distanceance is invalid",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: -10.0,
					DateTime: time.Date(2023, 6, 18, 10, 0, 0, 0, time.UTC),
				},
			},
			expectedErr: "invalid distance",
		},
		{
			name: "Should return -2 if date is invalid",
			input: []handleerrorsrefactoring.Segment{
				{
					Distance: 10.0,
					DateTime: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			expectedErr: "invalid date",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fare, err := handleerrorsrefactoring.CalculateRide(tt.input)
			if err != nil {
				assert.EqualError(t, err, tt.expectedErr)
			}
			assert.Equal(t, tt.expectedFare, fare)
		})
	}
}
